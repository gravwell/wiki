# 圧縮

ストレージはデフォルトで圧縮されており、ストレージ処理の負荷をCPUへ移すのに役立っています。Gravwellは、スケールワイドパラダイム最新のハードウェア向けに構築された高度な非同期システムです。最新のシステムは一般的にCPUに過剰なまでプロビジョニングされていますが、一方で大容量ストレージの処理能力はそれほどではありません。データを圧縮することで、Gravwellはストレージリンクのストレスを軽減しながら、データの圧縮と解凍を非同期的に行うため、余剰のCPUサイクルを活用することができます。その結果、低速の大容量記憶装置ででも、圧縮を採用すると検索が高速化されます。圧縮は各ウェルごとに独立して設定でき、ホットデータとコールドデータで異なる圧縮設定が可能です。

特筆すべき例外は、あまり圧縮できないデータです（まったく圧縮できない場合）。このような状況では、データを圧縮しようとすると、ストレージ容量や速度が実際に改善されることなく、CPU 時間が消費されます。それによく当てはまる例は、生のネットワーク・トラフィックで、暗号化と高いエントロピーが効果的な圧縮を妨げています。ウェルの圧縮を無効にするには、"Disable-Compression=true" ディレクティブを追加します。

## 圧縮の設定

Gravwellは、デフォルト圧縮と透過的圧縮の2種類の圧縮をサポートしています。デフォルトの圧縮は、[snappy](https://en.wikipedia.org/wiki/Snappy_%28compression%29)圧縮システムを使用して、ユーザー空間で圧縮と解凍を実行します。一方、デフォルトの圧縮システムはすべてのファイルシステムと互換性があります。透過的な圧縮システムは、基礎となるファイルシステムを使用して、透過的なブロックレベルの圧縮を提供します。

透過的圧縮は、圧縮されていないページキャッシュを維持して、ホストカーネルの圧縮/解凍作業負荷を低減することを可能にします。透過的圧縮は非常に高速で効率的な圧縮/解凍を可能にしますが、基礎となるファイルシステムが透過圧縮をサポートしている必要があります。現在、[BTRFS](https://btrfs.wiki.kernel.org/index.php/Main_Page) と [ZFS](https://wiki.archlinux.org/index.php/ZFS) ファイルシステムがサポートされています。

注意：透過的圧縮は、ストレージ全体に影響を及ぼすageoutルールに重要な意味を持ちます。詳細は [エイジアウトのドキュメント](ageout.md) を参照してください。

**Disable-Compression**
デフォルト値: `false`
例: `Disable-Compression=true`
ウェル全体の圧縮が無効になっているため、ホットストレージ、コールドストレージの両方の場所で圧縮が用いられません

**Disable-Hot-Compression**
デフォルト値: `false`
例: `Disable-Hot-Compression=true`
ホットストレージの場所の圧縮は無効になっています。

**Disable-Cold-Compression**
デフォルト値: `false`
例: `Disable-Cold-Compression=true`
コールドストレージの場所の圧縮を無効にします。コールドストレージの場所が指定されていない場合、この設定は何の効果もありません。

**Enable-Transparent-Compression**
デフォルト値: `false`
例: `Enable-Transparent-Compression=true`
Gravwellはストレージデータを圧縮可能なものとしてマークし、圧縮操作を実行するかどうかはカーネルに任せます。

**Enable-Hot-Transparent-Compression**
デフォルト値: `false`
例: `Enable-Hot-Transparent-Compression=true`
Gravwellはホットストレージのデータを圧縮可能なものとしてマークし、圧縮操作を実行するかどうかはカーネルに任せます。

**Enable-Cold-Transparent-Compression**
デフォルト値: `false`
例: `Enable-Cold-Transparent-Compression=true`
Gravwell はコールドストレージのデータを圧縮可能なものとしてマークし、圧縮操作を実行するかどうかはカーネルに任せます。

注釈: 透過圧縮が有効になっていて、一方で基礎となるファイルシステムが透過圧縮と互換性がないと検出された場合、実質的にはデータは非圧縮となり、Gravwellはユーザーに通知を送信します。

警告: ホットストレージとコールドストレージの場所が圧縮に関して互換性がない場合、Gravwellはデータをホットからコールドにエージアウトするための追加作業を実行する必要があります。加速設定が有効になっている場合、Gravwellはエイジアウトを実行する際にデータのインデックスを再作成します。そのため、互換性のない圧縮設定は、エイジアウト中に大きなオーバーヘッドを発生させる可能性があります。非圧縮データは透過的圧縮されたデータと互換性がありますが、デフォルト圧縮は非圧縮または透過的に圧縮されたデータと互換性がありません。Gravwellは互換性のない圧縮でも完全に機能しますが、端的に、エイジアウトの間、インデクサーにはより高い負荷で動作させることのなります。

## 圧縮とレプリケーション

[レプリケーションシステム](replication.md)は、通常のウェルストレージと同じルールに従います。レプリケーションシステムによるデータ複製（レプリケーション）は、透過的圧縮、デフォルト圧縮、または全くの非圧縮に設定することができます。ウェル内のホットストレージとコールドストレージで適用されている互換性と圧縮に関する同じルールが、レプリケーションされたデータとレプリケーションピアにも適用されます。レプリケーションピアが互換性のない形式の圧縮の設定をされている場合、障害後の復元時にインデクサーが実行する作業が大幅に増えます。最高のパフォーマンスを得るために、Gravwellでは、ホットストレージ、コールドストレージ、およびレプリケーションストレージとで同じ圧縮方式を使用することを推奨しています。

レプリケーションストレージの圧縮は、`Disable-Compression`と`Enable-Transparent-Compression`の2つのディレクティブによって制御されます。デフォルトの圧縮方式はsnappy圧縮方式です。

## 圧縮設定例

全てのウェルについて圧縮を無効にしたストレージウェル設定例:

```
[Storage-Well "network"]
	Location=/opt/gravwell/storage/network
	Cold-Location=/mnt/storage/gravwell_cold/network
	Tags=pcap
	Max-Hot-Data-GB=100
	Max-Cold-Data-GB=1000
	Delete-Frozen-Data=true
	Disable-Compression=true
```

ホットストレージの場所では圧縮を無効にし、コールドストレージの場所では透過圧縮を有効にしたストレージウェルの設定例。この設定では二つのストレージの圧縮方式には互換性ありと捉えられ、エイジアウト時に必要となる追加処理はありません。

```
[Storage-Well "syslog"]
	Location=/opt/gravwell/storage/syslog
	Cold-Location=/mnt/storage/gravwell_cold/syslog
	Tags=syslog
	Max-Hot-Data-GB=100
	Max-Cold-Data-GB=1000
	Delete-Frozen-Data=true
	Disable-Hot-Compression=true
	Enable-Cold-Transparent-Compression=true
```

ホットストレージの場所で透過的圧縮を有効にし、コールドウェルでデフォルトのユーザースペース圧縮を有効にしたストレージウェルの例。この設定は互換性なしと捉えられ、データのエイジアウト時に追加のオーバーヘッドが発生します。

```
[Storage-Well "windows"]
	Location=/opt/gravwell/storage/windows
	Cold-Location=/mnt/storage/gravwell_cold/windows
	Tags=windows
	Max-Hot-Data-GB=100
	Max-Cold-Data-GB=1000
	Delete-Frozen-Data=true
	Enable-Hot-Transparent-Compression=true
	Disable-Cold-Compression=true
```

レプリケーションストレージに透過的圧縮を使用するレプリケーション設定の例。

```
[Replication]
	Peer=indexer1
	Peer=indexer2
	Peer=indexer3
	Peer=indexer4
	Storage-Location=/mnt/storage/replication
	Enable-Transparent-Compression=true
```
