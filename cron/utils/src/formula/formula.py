#!/usr/bin/python3
import os
import inquirer


def Run(action, params):

    job, freq, weekday, month, timestamp = params
    
    if action == 'Create cron job':
        create_job(job, freq, weekday, month, timestamp)
    elif action == 'List cron jobs':
        list_jobs()
    elif action == 'Remove cron job':
        remove_job()

def get_weekday(weekday):
    day_as_int = {
        "sunday":       0,
        "monday":       1,
        "tuesday":      2,
        "wednesday":    3,
        "thursday":     4,
        "friday":       5,
        "saturday":     6
    }
    return day_as_int[weekday.lower()]

def escape_chars(string):
    '''
        escape common shell special characters
    '''
    string = string.replace("\\","\\\\\\")
    string = string.replace("*", "\*")
    string = string.replace('"', '\\"')
    
    return string

def get_cronjobs():
    '''
        fetch the cron jobs running in the machine
    '''
    output = os.popen("crontab -l")
    cron_jobs = []
    ''' get ouput cron jobs, one per line,
        and remove the empty line in the end of the output'''
    cron_jobs = output.read().split("\n")[:-1]

    if cron_jobs:
        #add the cancel option
        cron_jobs += ['Cancel']
    
    return cron_jobs

def create_job(job, freq, weekday, month, timestamp):
    '''
        create a new cron job with the ritchie inputs
    '''
    hour, minutes = timestamp.split(":")
    
    if freq != "Weekly":
        day = "*"
    else:
        day = get_weekday(weekday)
    
    if not month:
        month = '*'

    #normalize input to use always the same kind of string parameters, to prevent escape errors
    job = job.replace("'",'"').replace('"', '\\"')

    cmd = 'crontab -l | { cat; echo "%s %s %s * %s %s"; } | crontab - ;' % (minutes, hour, month, day, job)
    output = os.popen(cmd)

    if not output.read():  #if there is no output after execution, it means the program ran smoothly
        print("[+] - Cron Job created successfully")
    else:
        print("[-] - Something odd occured")
        print("[!] - Ouput:", output)

def list_jobs():
    '''
        display the cron jobs running in the machine
    '''
    cron_jobs = get_cronjobs()
    if cron_jobs:
        print("Running cron jobs:")
        for job in cron_jobs[:-1]:
            print(f"[+] - {job}")
    else:
        print("No cron jobs running")

def remove_job():
    '''
        unschedule the cron job selected on the available list
    '''
    cron_jobs = get_cronjobs()

    #only run if there is at least one cron job on the list
    if cron_jobs:
    
        select_cron_job = [inquirer.List('jobs', message="Which job you want to unschedule?", choices=cron_jobs),]
        
        answer = inquirer.prompt(select_cron_job)
        
        job = escape_chars(answer["jobs"])

        if job != 'Cancel':

            cmd = f'crontab -l | grep -v "{job}" | crontab -'
            confirm_action = [inquirer.Confirm('confirm', message=f"Do you really want to remove {answer['jobs']}?")]
            answer = inquirer.prompt(confirm_action)
            
            if answer['confirm']:
                output = os.popen(cmd)
                parsed_output = output.read().lower()
                if  parsed_output:
                    print("[-] something went wrong.")
                    print(parsed_output)
                else:
                    print("[+] - Cron Job unscheduled successfully!")

    else:
        print("[+] no jobs to unschedule")
