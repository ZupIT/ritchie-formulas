# -*- coding: utf-8 -*-

from colorama import Fore
from tabulate import tabulate

from .ec2 import Instances


def handle(params: dict):
    instance = Instances(params)
    content = instance.get_all()
    headers = [
        "Name",
        "Instance ID",
        "State",
        "Private IP",
        "Public IP",
        "Launch Time",
    ]
    rows = []
    for instance in content:
        color = get_status_color(instance["State"])
        state = f"{color}{instance['State']}{Fore.RESET}"
        rows.append(
            [
                str(instance["Name"]),
                str(instance["ID"]),
                str(state),
                str(instance["Private IP"]),
                str(instance["Public IP"]),
                str(instance["Launch Time"]),
            ]
        )

    table = tabulate(rows, headers, tablefmt="pretty")
    print(f"\n{table}\n")


def get_status_color(state: str) -> int:
    status = {
        "running": Fore.GREEN,
        "pending": Fore.YELLOW,
        "stopping": Fore.YELLOW,
        "stopped": Fore.RED,
        "terminated": Fore.RED,
        "shutting-down": Fore.YELLOW,
    }
    return status.get(state, Fore.WHITE)
