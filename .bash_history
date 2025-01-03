ls
df -ah
fdisk -l
sudo fdisk -l
mkdir /mnt/disk
sudo mkdir /mnt/disk
sudo mount /dev/sda3 /mnt/disk/
ls
cd ..
ls
ls -lah
cd ..
ls -lah
free -h
sudo apt install hddtemp lm-sensors
sudo apt update && sudo apt upgrade -y
sudo apt install hddtemp lm-sensors
sudo apt install lm-sensors
sensors
exit
ls
htop
nano /etc/ssh/sshd_config
sudo vim /etc/ssh/sshd_config
sudo service sshd restart
sudo service ssh restart
sudo service ssh status
ip a
sensors
cat /sys/class/power_supply/BAT*/power_now
cat /sys/devices/platform/smapi/BAT0/power_now
find /sys/ -type f -name power_now 2> /dev/null
sudo fdisk -l
df -ahl
vgdisplay
sudo vgdisplay
sudo shutdown -h now
ls
nano .ssh/authorized_keys 
htop
exit
htop
exit
sensors
exit
sudo apt install stress s-tui
nproc
stress –cpu 4 –timeout 60
stress –-cpu 4 –timeout 60
man stress
stress –c 4 –timeout 60
man stress
stress 
stress --cpu 4 --timeout 10s
stress –-cpu 4 –-timeout 60s
stress --cpu 4 --timeout 60s
stress --cpu 4 --timeout 1000s
sudo sensors-detect
pwmconfig
sudo apt install fancontrol
pwmconfig
sudo pwmconfig
sudo sensors-detect
sensors
sensors | grep fan
sensors 
sudo pwmconfig
sudo systemctl start fancontrol
sudo systemctl enable fancontrol
sensors 
sudo apt-get install macfanctld
sudo systemctl enable macfanctld
sudo systemctl start macfanctld
stress --cpu 4 --timeout 1000s
sudo nano /etc/macfanctl.conf
sudo cat /etc/macfanctl.conf
sudo nano /etc/macfanctl.conf
sudo systemctl restart macfanctld
sensors 
stress --cpu 4 --timeout 1000s
sensors 
stress --cpu 4 --timeout 1000s
sensors 
htop
sudo nano /etc/macfanctl.conf
sudo systemctl restart macfanctld
sudo nano /etc/macfanctl.conf
sudo systemctl restart macfanctld
sensors 
sensors
stress --cpu 4 --timeout 1000s
sensors
man scp
sensors
ip a
ls
rm -rf smCloud/
ls
df -h
ls
cd smCloud/
ls
cd backend/
ls
cd ..
rm -rf smCloud/
nano .ssh/authorized_keys 
cd smCloud/
nano .git/
cd .git/
ls
cd ..
df -hal
htop
sensors
exit
sensors
df -ahl
sensors
sensors 
htop
journalctl 
journalctl  -t
journalctl  -n 20
journalctl  -n 100
journalctl  scp
journalctl -u scp
journalctl -u ssh
cat /var/log/messages
cat  /proc/sys/kernel/panic
sudo journalctl | grep -i "shutdown"
sudo journalctl | grep -i "critical"
sudo journalctl | grep -i "power"
sudo dmesg
sudo dmesg | tail -n 100
sudo journalctl  | grep -i "out of memory"
htop
ls
htop
exit
sensors
exit
htop
ip a
ping 192.168.15.10
ping 192.168.15.97
nano /etc/default/grub
sudo nano /etc/default/grub
sudo update-grup2
sudo update-grub2
reboot -h now
sudo reboot -h now
ls
lspci -nn | grep -i network
sudo apt install bcmwl-kernel-source
apt-get update --fix-missing
sudo apt-get update --fix-missing
sudo apt install bcmwl-kernel-source
sudo apt update
sudo apt upgrade
ip a
sudo apt install bcmwl-kernel-source
ip a
sudo modprobe wl
ip a
sudo reboot
netplan
netplan generate
sudo netplan generate
nano /etc/netplan/50-cloud-init.yaml 
sudo vim /etc/netplan/50-cloud-init.yaml 
ip a
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml
ip a
ipa 
ip a
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml
ip a
vim /etc/netplan/50-cloud-init.yaml 
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml 
ip a
ls
ip a
wl
w
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo vim -X /etc/netplan/50-cloud-init.yaml 
exit
ip a
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml 
sudo vim /etc/netplan/50-cloud-init.yaml 
ip a
ipa 
ip a
sudo vim /etc/netplan/01-netcfg.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml 
sudo netplan apply 
ip a
reboot
xclock
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo iwlist wlp2s0 scan
sudo iwlist wlp2s0 scan > wifis
nano wifis 
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml 
ip a
sudo vim /etc/netplan/50-cloud-init.yaml 
sudo netplan apply /etc/netplan/50-cloud-init.yaml 
ip a
sudo netplan try
ip a
sudo netplan apply
sudo mv /etc/netplan/50-cloud-init.yaml /etc/netplan/01-netcfg.yaml
sudo netplan apply
ip a
nano /etc/netplan/01-netcfg.yaml 
sudo nano /etc/netplan/01-netcfg.yaml 
sudo netplay apply
sudo netplan apply
ip a
sudo netplan /etc/netplan/01-netcfg.yaml 
sudo netplan apply /etc/netplan/01-netcfg.yaml 
ip a
sudo netplan apply /etc/netplan/01-netcfg.yaml --debug 
sudo netplan --debug apply /etc/netplan/01-netcfg.yaml
sudo netplan --debug apply /etc/netplan/01-netcfg.yaml > debug
sudo netplan --debug apply /etc/netplan/01-netcfg.yaml &> debug
nano debug 
scp debug thiagobaldino@ 192.168.15.10:/Users/thiagobaldino
scp debug thiagobaldino@192.168.15.10:/Users/thiagobaldino
ip a
sudo netplan apply
nano /etc/netplan/01-netcfg.yaml 
sudo nano /etc/netplan/01-netcfg.yaml 
sudo netplan apply
ip a 
sudo netplan apply --debug
sudo netplan --debug apply 
ip a
sudo nano /etc/netplan/01-netcfg.yaml 
sudo netplan try
ls
ip 
sudo netplan --debug apply
ip a
systemctl status netplan-wpa-wlp2s0.service
sudo systemctl restart netplan-wpa-wlp2s0.service
systemctl status netplan-wpa-wlp2s0.service
journalctl -u  netplan-wpa-wlp2s0.service
which wpa_supplicant
sudo apt install wpasupplicant
systemctl status netplan-wpa-wlp2s0.service
sudo systemctl restart netplan-wpa-wlp2s0.service
systemctl status netplan-wpa-wlp2s0.service
ip a
htop
sensors
htop
sensors
df -ahl
vgdisplay
sudo vgdisplay
lvextend -l +100%FREE /dev/mapper/ubuntu--vg-ubuntu--lv
fdisk -l
sudo fdisk -l
sudo fdisk /dev/sdb
sudo fdisk /dev/sdb1
sudo fdisk /dev/sdb2
sudo fdisk /dev/sdb
mkfs.ext4 /dev/sdb1
sudo mkfs.ext4 /dev/sdb1
ls
df -ahl
ls
cd /
sudo mkdir /mnt/temp
sudo mount /dev/sdb1 /mnt/temp/
rsync -av /mnt/temp/ /home/thiago/
sudo rsync -av /mnt/temp/ /home/thiago/
sudo rsync -av /mnt/temp/ /home/
sudo rsync -av /mnt/temp/ /home
sudo rsync -av /home /mnt/temp/
ls
cd /mnt/temp/
ls
cd home/
ls
cd thiago/
ls
cd ..
cd /mnt/temp/
ls
cd home/
ls
mv thiago/ ../
sudo mv thiago/ ../
cd ..
ls
cd home/
ls
cd ..
rm -rf home/
sudo rm -rf home/
cd ..
sudo umount /mnt/temp 
sudo rm -rf /home/*
sudo mount /dev/sdb1 /home/
cd /home/
ls
sudo blkid /dev/sdb1
sudo vim /etc/fstab 
sudo reboot
exit
df -hal
lvextend -l +100%FREE /dev/mapper/ubuntu--vg-ubuntu--lv
sudo lvextend -l +100%FREE /dev/mapper/ubuntu--vg-ubuntu--lv
sudo resize2fs /dev/mapper/ubuntu--vg-ubuntu--lv
df -hal
ls
sudo apt install golang-go
go start
go run
ls
cd smCloud/
ls
cd backend/
ls
cd ..
exit
sudo apt install unzip
pip3 install --upgrade pynvim
pip install --upgrade pynvim
sudo apt install python3-pip
pip install --upgrade pynvim
pip3 install --upgrade pynvim
sudo pip3 install --upgrade pynvim
sudo pip3 install --upgrade python3-pynvim
sudo apt install python3-pynvim
apt install fd
sudo apt install fd
sudo apt install fd-find
sudo apt-get install ripgrep
unzip
q
sudo apt install xclip
cd .config/
ls
cd htop/
ls
cd ..
ls
cd ..
ls
cd ..
ls
cd .config
cd thiago/
ls
cd .config/
ls
nvim 
udo apt  install neovim
sudo apt  install neovim
nvim 
cd nvim/
ls
ls init.lua
cat init.lua
curl -fLo ~/.vim/autoload/plug.vim --create-dirs     https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
cd ..
chown -r thiago .config/
chown -R thiago .config/
curl -fLo ~/.vim/autoload/plug.vim --create-dirs     https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
       https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'
nvim
ls
cd .config/
ls -lah
cd nvim/
ls -lah
cd ..
chmod ugo+rwx nvim
ls -lah
cd ..
nvim .
cd ..
ls -lah
sudo chown thiago thiago/
ls -lah
cd thiago/
ls
nvim 
which nvim
cd /usr/bin
ls
ls | grep nvim
ls -ahl | grep nvim
cd
wget https://github.com/neovim/neovim/releases/download/v0.10.2/nvim-macos-x86_64.tar.gz
tar -xvf nvim-macos-x86_64.tar.gz 
ls
crm -rf nvim-macos-x86_64*
rm -rf nvim-macos-x86_64*
wget https://github.com/neovim/neovim/releases/download/v0.10.2/nvim-linux64.tar.gz
tar -xvf nvim-linux64.tar.gz 
ls
cd nvim-linux64/
ls
cd bin/
ls
ls -lah
./nvim 
./nvim --version
which nvim
cd ..
ls
cd lib/
ls
cd nvim/
ls
cd parser/
ls
cat vim.so 
clear
cd /usr/lib/nvim/
ls
cd parser/
ls
cd ..
sudo apt remove neovim 
mv ~/nvim-linux64/lib/nvim/ ./
sudo mv ~/nvim-linux64/lib/nvim/ ./
ls
cd ..
cd share/
ls
sudo mv ~/nvim-linux64/share/nvim/ ./
rm -rf nvim/
sudo rm -rf nvim/
sudo mv ~/nvim-linux64/share/nvim/ ./
cd ..
cd bin/
sudo mv ~/nvim-linux64/bin/nvim ./
nvim --version
nvim .
nvim
cd
nvim
exity
exit
nvim 
rm -rf nvim-linux64*
ls
rm -4f debug wifis 
rm -rf debug wifis 
ls -lah
cat .sudo_as_admin_successful 
nvim
modprobe kvm_intel
kvm-ok
lsmod | grep kvm
ls -al /dev/kvm
sudo usermod -aG kvm thiago
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
ls
echo   "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable"
echo   "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" |   sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
docker run hello-world
sudo usermod -aG docker thiago
exit
htop
exit
docker run hello-world
docker ls 
docker ps
docker ps -a
docker image rm hellow-world
docker image rm hello-world
docker rm ed58a1be2829
docker image rm hello-world
df -lah
sensors
passwd thiago 
exit
sudo a
ip a
fdisk -l
sudo fdisk -l
cd /mnt
ls
sudo mount /dev/sdc1 disk/
sudo fdisk -l
sudo mount /dev/sdc1 disk/
cd disk/
ls
cat badblocks.log 
ls
df -hal
sudo fdisk -l
umount /dev/sdc1
sudo umount /dev/sdc1
cd ..
sudo umount /dev/sdc1
sudo fdisk -l
exit
