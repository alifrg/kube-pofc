# -*- mode: ruby -*-
# vi: set ft=ruby :

# Global Variables
K8S_NAME = "tnt-kube"
HA_NUM = 2
MASTERS_NUM = 3
WORKERS_NUM = 2
NODES_NUM = 3
IP_BASE = "192.168.30."

Vagrant.configure("2") do |config|
  # Global Configs
  config.vm.box = "ubuntu/bionic64"

  # BootStrap Loadbalancers
  (1..HA_NUM).each do |i|
    config.vm.define "ha#{i}" do |ha|
      ha.vm.hostname = "ha#{i}.kube.tnt"

      ha.vm.provider "virtualbox" do |vb|
        # Customize the amount of memory and cpu on the VMs:
        vb.memory = 256
        vb.cpus = 1
      end
      #  Set IP address for masters
      ha.vm.network "private_network", ip: "#{IP_BASE}#{i + 10}"
      # Provision the loadbalancer using Ansible When all nodes are created
      if i == HA_NUM
      ha.vm.provision :ansible do |ansible|
          # Disable default limit to connect to all the machines
          ansible.limit = "kubebalancer_nodes"
          ansible.playbook = "kubeadm.yml"
          ansible.become = true
          ansible.groups = {
            "kubebalancer_nodes" => ["ha[1:#{i}]"],

          }
          ansible.extra_vars = {
                    host_os: "ubuntu",
                    base_ip: "#{IP_BASE}",
                    ns1_ip: "8.8.8.8",
                    ns2_ip: "1.1.1.1",
                    LOAD_BALANCER_DNS: "controller.kube.tnt",
                    LOAD_BALANCER_PORT: "6443",
                    CONTROLLER_1: "#{IP_BASE}21",
                    CONTROLLER_2: "#{IP_BASE}22",
                    CONTROLLER_3: "#{IP_BASE}23",
                    keepalived_auth_pass: "5dcc24e828a39963fac4208f",
                    keepalived_role: "MASTER",
                    keepalived_router_id: "52",
                    keepalived_shared_ip: "#{IP_BASE}13",
                    keepalived_check_process: "keepalived",
                    keepalived_priority: "100",
                    keepalived_backup_priority: "50"
          }
      end
      end
      end
  end

  # BootStrap Master Nodes
  (1..MASTERS_NUM).each do |i|
      config.vm.define "master#{i}" do |master|
        master.vm.hostname = "master#{i}.kube.tnt"

        master.vm.provider "virtualbox" do |vb|
          # Customize the amount of memory and cpu on the VMs:
          vb.memory = 1024
          vb.cpus = 2
        end
        #  Set IP address for masters
        master.vm.network "private_network", ip: "#{IP_BASE}#{i + 20}"
      # Provision the Kube Masters using Ansible When all nodes are created
      if i == MASTERS_NUM
        master.vm.provision :ansible do |ansible|
            ansible.limit = "kubemaster_nodes"
            ansible.playbook = "kubeadm.yml"
            ansible.become = true
            ansible.groups = {
              "docker_nodes" => ["master[1:#{i}]"],
              "kubemaster_nodes" => ["master[1:#{i}]"],
              "kubeadm_nodes" => ["master[1:#{i}]"],

            }
            # Set INIT_MODE Variable which determinse the node for running kubeadm init
            # If you need more than 3 masters you should define them here as well
            ansible.host_vars = {
              "master1" => {
                "INIT_NODE" => TRUE,
                "master1_ip" => "#{IP_BASE}21",
                "pod_network_cidr" => "192.168.112.0/20",
              },
              "master2" => {
                "INIT_NODE" => FALSE,
              },
              "master3" => {
                "INIT_NODE" => FALSE,
              },
            }
            ansible.extra_vars = {
                      host_os: "ubuntu",
                      base_ip: "#{IP_BASE}",
                      ns1_ip:  "192.168.30.11",
                      ns2_ip:  "192.168.30.12",
                      master_node: "true",
                      LOAD_BALANCER_DNS: "controller.kube.tnt",
                      LOAD_BALANCER_PORT: "6443",
                      CONTROLLER_1: "#{IP_BASE}21",
                      CONTROLLER_2: "#{IP_BASE}22",
                      CONTROLLER_3: "#{IP_BASE}23",
            }
        end
      end
      end
    end

# BootStrap Worker Nodes
(1..WORKERS_NUM).each do |i|
  config.vm.define "worker#{i}" do |master|
    master.vm.hostname = "worker#{i}.kube.tnt"

    master.vm.provider "virtualbox" do |vb|
      # Customize the amount of memory and cpu on the VMs:
      vb.memory = 1024
      vb.cpus = 2
    end
    #  Set IP address for masters
    master.vm.network "private_network", ip: "#{IP_BASE}#{i + 30}"
  # Provision the Kube Masters using Ansible When all nodes are created
  if i == WORKERS_NUM
    master.vm.provision :ansible do |ansible|
        ansible.limit = "kubeworker_nodes"
        ansible.playbook = "kubeadm.yml"
        ansible.become = true
        ansible.groups = {
          "docker_nodes" => ["worker[1:#{i}]"],
          "kubeworker_nodes" => ["worker[1:#{i}]"],
          "kubeadm_nodes" => ["worker[1:#{i}]"],

        }
        # ansible.host_vars = {
        #   "worker1" => {
        #     "node_ip" => "#{IP_BASE}31",
        #   },
        #   "worker" => {
        #     "node_ip" => "#{IP_BASE}32",
        #   },
        #   # "worker3" => {
        #   #   "node_ip" => "#{IP_BASE}31",
        #   # },
        # }
        ansible.extra_vars = {
                  host_os: "ubuntu",
                  base_ip: "#{IP_BASE}",
                  ns1_ip:  "192.168.30.11",
                  ns2_ip:  "192.168.30.12",
                  LOAD_BALANCER_DNS: "controller.kube.tnt",
                  LOAD_BALANCER_PORT: "6443",
                  CONTROLLER_1: "#{IP_BASE}21",
                  CONTROLLER_2: "#{IP_BASE}22",
                  CONTROLLER_3: "#{IP_BASE}23",
        }
    end
  end
  end
end
end