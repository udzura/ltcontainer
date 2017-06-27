# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"
  config.vm.provision "shell", inline: (<<-SHELL).gsub(/^ {4}/, '')
    if ! docker run hello-world; then
      apt-get -y install \
          apt-transport-https \
          ca-certificates \
          curl \
          software-properties-common
      curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
      add-apt-repository -y \
         "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
         $(lsb_release -cs) \
         stable"
      apt-get -y update
      apt-get -y install docker-ce
    fi
  SHELL
end
