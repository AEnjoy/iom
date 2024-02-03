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

def install_nodejs():
    nodejs_version = "v20.11.0"
    nodejs_tar_file = f"node-{nodejs_version}-linux-x64.tar.xz"
    nodejs_url = f"https://nodejs.org/dist/{nodejs_version}/{nodejs_tar_file}"
    os.system(f"wget {nodejs_url}")
    os.system(f"tar -xf {nodejs_tar_file}")
    os.system(f"mv node-{nodejs_version}-linux-x64/bin/node /usr/local/bin/")
    os.system(f"mv node-{nodejs_version}-linux-x64/bin/npm /usr/local/bin/")
    os.system(f"rm -rf node-{nodejs_version}-linux-x64")
    os.system(f"rm {nodejs_tar_file}") 
    os.system("npm install -g cnpm --registry=https://registry.npmmirror.com")
    os.system("npm install -g yarn --registry=https://registry.npmmirror.com")
    
def command_exists(command):
    try:
        return_code = subprocess.call(["which", command], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        return return_code == 0
    except Exception as e:
        print(f"Error: {e}")
        return False


def main():
    check_sudo()
    if not command_exists("node"):
        install_nodejs()
    if command_exists("cnpm"):
        subprocess.run(["cnpm", "install"])
    elif command_exists("yarn"):
        subprocess.run(["yarn", "install"])
    else:
        subprocess.run(["npm", "install"])
    subprocess.run(["npm", "run", "build"])
    
    #os.system("mv assets res")
    #os.system("mv favicon.ico ./res")
    #os.system("cp -r dist/* ../server")
   
    #subprocess.run(["docker", "build", "-t", "iom-dashboard:latest", "."])


if __name__ == '__main__':
    if platform.platform().startswith("Linux"):
        main()
    else:
        print("此脚本只能在Linux系统上运行。")
