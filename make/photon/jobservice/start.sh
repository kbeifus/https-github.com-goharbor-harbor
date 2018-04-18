#!/bin/sh
if [ -d /var/log/jobs ]; then
    chown -R 10000:10000 /var/log/jobs/
fi
if [ -d /var/log/jobs/scan_job ]; then
    chmod +x /var/log/jobs/scan_job
fi

sudo -E -u \#10000 "/harbor/harbor_jobservice" "-c" "/etc/jobservice/config.yml"

