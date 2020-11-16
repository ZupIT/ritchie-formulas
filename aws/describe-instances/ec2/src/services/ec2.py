# -*- coding: utf-8 -*-

import boto3


class Instances:
    def __init__(self, config: dict):
        self.ec2 = boto3.resource(
            service_name="ec2",
            region_name=config["region"],
            aws_access_key_id=config["access_key"],
            aws_secret_access_key=config["secret_key"],
        )
        self.__set_filters(config.get("instance_state", ""))

    def __set_filters(self, config: str):
        self.filters = [{}]
        if config == "running":
            filter = {"Name": "instance-state-name", "Values": ["running"]}
            self.filters = [filter]
        elif config == "stopped":
            filter = {"Name": "instance-state-name", "Values": ["stopped"]}
            self.filters = [filter]

    def parse_data(self, instances):
        ec2info = []
        for instance in instances:

            if instance.tags is None:
                continue

            for tag in instance.tags:
                if "Name" in tag["Key"]:
                    name = tag["Value"]
                    break

            ec2info.append(
                {
                    "Name": name,
                    "ID": instance.id,
                    "Type": instance.instance_type,
                    "State": instance.state["Name"],
                    "Private IP": instance.private_ip_address,
                    "Public IP": instance.public_ip_address,
                    "Launch Time": instance.launch_time,
                }
            )
        return ec2info

    def get_all(self) -> dict:
        """Get all instances.

        Returns:
            dict: return dict with all instances in aws.
        """
        instances = self.ec2.instances.filter(Filters=self.filters)
        parsed_instances = self.parse_data(instances)
        return parsed_instances

    def get_instance(self, filter: list):
        instance = None
        instances = self.ec2.instances.filter(Filters=filter)

        for item in instances:
            instance = item
        return instance

    def get_instance_by_name(self, instance_name: str):
        filter = [{"Name": "tag:Name", "Values": [instance_name]}]
        return self.get_instance(filter)

    def get_instance_by_id(self, instance_id: str):
        filter = [{"Name": "instance-id", "Values": [instance_id]}]
        return self.get_instance(filter)
