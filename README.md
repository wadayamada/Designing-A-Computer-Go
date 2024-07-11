# CPUのエミュレーターをGoで作る
命令セットは以下
## 命令セット(拡張TD4)
|  opecode  |  アセンブラ  | 概要　|
| ---- | ---- |----|
|  0001  |  MOV A, B  | A <- B|
|  0100  |  MOV B, A  | B <- A|
|  0011  |  MOV A, IMM | A <- IMM|
|  0111  |  MOV B, IMM | B <- IMM|
|  0010  |  IN A  | A <- IN_A|
|  0110  |  IN B  | B <- IN_B|
|  1001  |  OUT B  | OUT <- B|
|  1011  |  OUT IMM  |OUT <- IMM|
|  0000  |  ADD A, IMM  |A <- A+IMM|
|  0101  |  ADD B, IMM  |B <- B+IMM|
|  1000  |  ADD A, B | A <- A+B|
|  1111  |  JMP IMM  | IP <- IMM|
|  1110  |  JNC IMM  | if CF==0: IP <-IMM|
|  1010  |  MOV C, A | C <- A|
|  1100  |  MOV A, C | A <- C|
|  1101  | 空いてる ||
|  その他  |  未定義 |　使用禁止|

## チューリング完全
チューリングマシンは特定の計算を行う計算モデルのこと。入力部に入力された二つの数字の和を計算するなど。
万能チューリングマシンは渡された任意のチューリングマシンの動作を再現できる計算モデルのこと。
足し算のチューリングマシンを渡されたら足し算ができるし、掛け算のチューリングマシンを渡されたら掛け算ができる。

万能チューリングマシンを模倣できる計算モデルはチューリング完全というらしい
NANDはチューリング完全らしい

https://qiita.com/payanotty/items/ee0bd383cf9fdd335139

NANDでCPUのエミュレーターを作っていく

## NANDでNOT, AND, ORを作る

### NOT
NANDの両方に入力Aを入れれば良い
### AND
NOT(NAND)
### OR
ド・モルガンの法則

## Goについて
### 基本的な書き方
文法
https://zenn.dev/ak/articles/1fb628d82ed79b
命名
https://zenn.dev/keitakn/articles/go-naming-rules

### module, packageについて理解する
moduleは大きめな括り
api, shopcmsくらいの範囲なのかな？
今回はcomputerで切った

packageはその中の括り
今回はlogicgate, nandで切った

https://zenn.dev/masaruxstudy/articles/7965c98289caf5

### 単体テストも実装する
テストテンプレートを自動生成できて便利
https://zenn.dev/nishisuke/articles/go-unit-test-with-code-generation

## multi plexer
入力から値を1つ選択するやつ

## やりたいこと
- チューリング完全であることを証明する
- 既存の命令セットと比較をする
- 算術演算を実装する
  - ~足し算~
  - ~引き算~
  - ~掛け算~
  - 割り算
- 解説ややったことをまとめる
- 計算可能性の観点において、チューリングマシンより計算能力のある計算モデルはないんじゃなかったっけ？だからチューリング完全であるBrainF\*ckでは今回作ったCPU作れるし、今回作ったCPUでもBrainF\*ck作れそう

## 時間があればやりたい
- godoc書きたい
