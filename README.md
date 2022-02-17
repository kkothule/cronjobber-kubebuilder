# Cronjobber

Cronjobber is re-written using kubebuilder with reference from project https://github.com/hiddeco/cronjobber. the cronjob controller from Kubernetes patched with time zone support.
This is only for pratice purpose

## Installation

```sh
# Install everything in one go
$ kubectl apply -f https://raw.githubusercontent.com/kkothule/cronjobber/main/deploy/manifests.yaml
```

## Usage

Instead of creating a [`CronJob`](https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/)
like you normally would, you create a `TZCronJob`, which works exactly
the same but supports an additional field: `.spec.timezone`. Set this
to the time zone you wish to schedule your jobs in and Cronjobber will
take care of the rest.

```yaml
apiVersion: cronjobber.github.com/v1
kind: TZCronJob
metadata:
  name: hello
spec:
  schedule: "*/1 * * * *"
  timezone: "Europe/Amsterdam"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: hello
              image: busybox
              args:
                - /bin/sh
                - -c
                - date; echo "Hello, World!"
          restartPolicy: OnFailure
```

## Credits

This application is derived from open source components. You can find
the original source code of these components below.

* https://github.com/hiddeco/cronjobber
* https://book.kubebuilder.io/quick-start.html
* [Kubernetes CronJob controller](https://github.com/kubernetes/kubernetes/tree/v1.13.3/pkg/controller/cronjob)
* [kubernetes/kubernetes#47266](https://github.com/kubernetes/kubernetes/pull/47266) by Adam Sunderland
