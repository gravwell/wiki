## Subnet

サブネットは、IPアドレスからサブネットを抽出するように設計されています。  これは、特定のサブネット内にのみ含まれる値を確認したり、発信元に基づいて攻撃者を分類したりするのに役立ちます。  デフォルトでは、サブネットモジュールはIPv4を想定していますが、`-6`フラグを介してIPv6を完全にサポートしています。  サブネットはデフォルトで'subnet'という名前の列挙値に抽出されますが、"as"フラグを使用すると、別のターゲットを指定できます。

### サポートされているオプション

* `-4`: IPv4サブネットとIPを探す
* `-6`: IPv6のサブネットとIPを探す

### 使用例

発信元サブネット別の失敗したSSHログイン試行のグラフ化

```
tag=syslog grep sshd | grep "Failed password for" | regex "\sfrom\s(?P<ip>\S+)\s" | subnet ip /16 | count by subnet | chart count by subnet limit 64
```

失敗したSSHログイン試行を特定のサブネットからのものだけにフィルタリング

```
tag=syslog grep sshd | grep "Failed password for" | regex "\sfrom\s(?P<ip>\S+)\s" | subnet ip /16 as attackersub | grep -e attackersub “34.22.1.0” | count by ip | sort by count desc | table ip count
```

パケットからソースおよび宛先サブネットを取得:

```
tag=pcap packet ipv4.SrcIP ipv4.DstIP | subnet SrcIP /16 as srcsub DstIP /16 as dstsub | table
```

![](subnet.png)