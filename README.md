# FinYAML

FinYAML is a command line tool to inject a File into a YAML template. It looks in the input template for a filename surrounded by `{{ }}` e.g {{"filename.txt"}} and replaces the tag with the contents of the file escaped for YAML

## Installation
installs the binary as FinYAML (case sensitive).
```shell 
go get github.com/reechar-goog/FinYAML
```
 This will install the binary as all lowercase if preferred:
```shell
go get github.com/reechar-goog/finyaml
```

## Usage
```shell
FinYAML sample-template.yaml
```
by default it ouputs to output.yaml. However you can specify the outputfile path
```shell
FinYAML -out=rendered-template.yaml sample-template.yaml
```

## Sample
sample-template.yaml:
```yaml
# Example of the Instance Template template usage.
#
# In this example, an instance template with Nginx is created.

imports:
  - path: templates/instance_template/instance_template.py
    name: instance_template.py

resources:
  - name: ci-template-string
    type: instance_template.py
    properties:
      diskImage: projects/cos-cloud/global/images/cos-stable-70-11021-51-0
      network: default
      machineType: f1-micro
      scheduling:
        preemptible: true
      metadata:
        items:
          - key: user-data
            value: {{"cloud_init.yaml"}}
```

out.yaml:
```yaml
# Example of the Instance Template template usage.
#
# In this example, an instance template with Nginx is created.

imports:
  - path: templates/instance_template/instance_template.py
    name: instance_template.py

resources:
  - name: ci-template-string
    type: instance_template.py
    properties:
      diskImage: projects/cos-cloud/global/images/cos-stable-70-11021-51-0
      network: default
      machineType: f1-micro
      scheduling:
        preemptible: true
      metadata:
        items:
          - key: user-data
            value: "#cloud-config\n\nusers:\n- name: cloudservice\n  uid: 2000\n\nwrite_files:\n- path: /etc/systemd/system/cloudservice.service\n  permissions: 0644\n  owner: root\n  content: |\n    [Unit]\n    Description=Start a simple docker container\n\n    [Service]\n    ExecStart=/usr/bin/docker run --rm -u 2000 --name=mycloudservice busybox:latest /bin/sleep 3600\n    ExecStop=/usr/bin/docker stop mycloudservice\n    ExecStopPost=/usr/bin/docker rm mycloudservice\n\nruncmd:\n- systemctl daemon-reload\n- systemctl start cloudservice.service\n\n# Optional once-per-boot setup. For example: mounting a PD.\nbootcmd:\n- fsck.ext4 -tvy /dev/[DEVICE_ID]\n- mkdir -p /mnt/disks/[MNT_DIR]\n- mount -t ext4 -O ... /dev/[DEVICE_ID] /mnt/disks/[MNT_DIR]"
```
