#! /usr/bin/env python3
# -*- coding: utf-8 -*-ff
import os
import sys
import platform

try:
    import subprocess
except ImportError:
    import pip

    pip.main(['install', 'subprocess'])
    import subprocess


def check_sudo():
    if os.geteuid() != 0:
        print("此脚本需要以root权限运行。")
        sys.exit(1)


def install_docker():
    print("正在安装Docker...")
    subprocess.run("curl -fsSL https://get.docker.com | sh", shell=True, check=True)
    subprocess.run("systemctl start docker", shell=True, check=True)
    subprocess.run("systemctl enable docker", shell=True, check=True)


def install_golang():
    subprocess.run(["wget", "https://golang.google.cn/dl/go1.21.6.linux-amd64.tar.gz"])
    subprocess.run(["tar", "-zxf", "go1.21.6.linux-amd64.tar.gz", "-C", "/usr/local"])
    with open(os.path.expanduser("~/.bashrc"), "a") as f:
        f.write("export PATH=$PATH:/usr/local/go/bin\n")
    subprocess.run(["source", "~/.bashrc"])
    subprocess.run(["go", "env", "-w", "GOPROXY=https://goproxy.cn"])


def command_exists(command):
    try:
        return_code = subprocess.call(["which", command], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        return return_code == 0
    except Exception as e:
        print(f"Error: {e}")
        return False


def main():
    check_sudo()
    if not command_exists("docker"):
        install_docker()
    if not command_exists("go"):
        install_golang()
    subprocess.run(["go", "mod", "tidy", "-y"])
    subprocess.run(["go", "build", "-o", "./main"])
    subprocess.run(["docker", "build", "-t", "iom-server:latest", "."])


if __name__ == '__main__':
    if platform.platform().startswith("Linux"):
        main()
    else:
        print("此脚本只能在Linux系统上运行。")
