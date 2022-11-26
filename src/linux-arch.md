# kde部分应用中文输入
使用 Fcitx 之前，您必须先设置一些环境设定变量：

如果您用 KDM, GDM, LightDM 等显示管理器，请在 ~/.xprofile 中加入以下代码；如果您用 startx 或者 Slim 启动，即使用 .xinitrc 的场合，则改在 ~/.xinitrc 中加入：

```sh
export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS="@im=fcitx" 
```

http://newstart.farbox.com/post/articles/fcitx


# mysql

https://wiki.archlinux.org/index.php/MariaDB


已经设置了一个用户kk
```
MariaDB [xxx]> GRANT ALL ON xxx.* TO 'kk'@'127.0.0.1' IDENTIFIED BY 'sss' WITH GRANT OPTION;
```

# kde shortcut

`C-:`  choise pasteboard history 


# wireguard

配置文件放在 `/etc/wireguard/wg0.conf` 里面
使用 `wg-quick up wg0` 来启用 Interface, 使用 `wg-quick down wg0` 来关闭。

使用 `systemctl enable wg-quick@wg0` 来自动启动。 

[arch-wiki](https://wiki.archlinux.org/index.php/WireGuard_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87))


# route: how to use two card (not worked)

```bash
route add -net 192.168.62.0 netmask 255.255.255.0 gw 192.168.1.1


route add -net 10.10.40.0 netmask 255.255.255.0 gw 10.10.30.1

http://10.10.40.11

10.10.30.164
wlp0s20f0u11

ip route add 10.10.40.0/24 via 10.10.30.164 dev wlp0s20f0u11



# ok
~/w/tmp [2]> sudo ip route del default via 192.168.0.1
[sudo] xx 的密码：
~/w/tmp> ping 10.10.40.11
PING 10.10.40.11 (10.10.40.11) 56(84) bytes of data.


sudo ip route add 10.10.40.0/24 via 10.10.30.254 dev wlp0s20f0u11

```

# yay cache

err: HTTP server doesn't seem to support byte ranges. Cannot resume
忽略缓存下载: 在询问clean build的时候选择ALL
```
==> Packages to cleanBuild?
==> [N]one [A]ll [Ab]ort [I]nstalled [No]tInstalled or (1 2 3, 1-3, ^4)
==>

```

# firefox
## change scroll bar to left

1. input`about:config` in address bar
2. search and change `layout.scrollbar.side` to 3
3. restart(must?)

# vnc

black screen use the default config

https://www.reddit.com/r/ManjaroLinux/comments/g7vs5i/black_screen_when_i_vnc_into_manjaro_guest_kde/
Deleted everything in the `~/.vnc/xstartup` file and added
`dbus-launch startplasma-x11`

```sh
vncserver -geometry 1600x1200 -randr 1600x1200,1440x900,1024x768
xrandr
xrandr -s 800x600
```