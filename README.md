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
NANDはチューリング完全らしい(NANDがあればチューリング完全な計算機が作れるって意味だと思う)

https://qiita.com/payanotty/items/ee0bd383cf9fdd335139

NANDでCPUのエミュレーターを作っていく

## 今回作るCPUがチューリング完全であることを証明する
brainfu\*ckはチューリング完全であるらしいので、今回のCPUでbrainfu\*kが作れることを証明する

### brainfu\*kで使える命令は以下の8つ
- `>`: pointerの位置をincrement
- `<`: pointerの位置をdecrement
- `+`: pointerの指す値をincrement
- `-`: pointerの指す値をdecrement
- `.`: pointerの指す値を出力する
- `,`: 入力から1バイト読み込んで、pointerが指す場所に格納する
- `[`: pointerの指す値が0なら、対応する]の直後にジャンプする
- `]`: pointerの指す値が0じゃないなら、対応する[の直後にジャンプする

### 実装する
- pointerの値はregisterAに持つ
  - `>`: pointerの位置をincrement
    - ADD A, 1
  - `<`: pointerの位置をdecrement
    - ADD A, 15
- MOV B, RAM[A]とMOV RAM[A], Bを追加する
  - `+`: pointerの指す値をincrement
    - MOV B, RAM[A]  
    - ADD B, 1
    - MOV RAM[A], B 
  - `-`: pointerの指す値をdecrement
    - MOV B, RAM[A]  
    - ADD B, 15
    - MOV RAM[A], B 
  - `.`: pointerの指す値を出力する
    - MOV B, RAM[A]  
    - OUT B
  - `,`: 入力から1バイト読み込んで、pointerが指す場所に格納する
    - IN B
    - MOV RAM[A], B 
- コンパイルの詳細は割愛する
  - `[`: pointerの指す値が0なら、対応する]の直後にジャンプする
    - コンパイルしたら、"pointerの指す値が0だったら、IMMのアドレスにジャンプするという処理"に帰着できそう
      - pointerの指す値が0だったらIMMにジャンプする処理
        - 15を足してCF=0だったら0なので、IMMにジャンプすれば良い
          - 0: MOV B, RAM[A]
          - 1: ADD B, 15
          - 2: JNC IMM
  - `]`
    - コンパイルしたら、"pointerの指す値が0じゃないなら、IMMのアドレスにジャンプする処理"に帰着できそう
      - pointerの指す値が0じゃないなら、IMMにジャンプする処理
        - 15を足して、CF=1だったら0じゃないので、IMMにジャンプすれば良い
          - 0: MOV B, RAM[A]
          - 1: ADD B, 15
          - 2: JNC 5 (15を足してCF=0だったら、0ということなので、次の命令に飛ぶ)
          - 3: ADD B, 0
          - 4: JNC IMM (15を足してCF=1だったら、0じゃないということで、3, 4の処理が実行され、CF=0なので、IMMに飛ぶ)

## チューリングマシンと計算可能性
> チューリングマシンは非常に強力な計算モデルである。チューリングマシンの定義を修正してより強力なモデルを作ろうとしても失敗する。実際、チャーチ＝チューリングのテーゼでは、チューリングマシンで判定できない言語を判定可能な計算モデルはないと推定されている。
https://ja.wikipedia.org/wiki/%E8%A8%88%E7%AE%97%E5%8F%AF%E8%83%BD%E6%80%A7%E7%90%86%E8%AB%96

brainfu\*kはチューリング完全だから、チューリングマシンと同じ計算能力を持つ。

拡張TD4でbrainfu\*kを作れたから、拡張TD4もチューリング完全。

チューリング完全なものはお互いでお互いを作れる。brainfu\*kでも拡張TD4を作れる。

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
- ~チューリング完全であることを証明する~
- 既存の命令セットと比較をする
  - 有名な命令セット
    - arm
    - risc
    - x86
    - x86-64
    - jvm
    - risc-v
    - ref: https://ja.wikipedia.org/wiki/%E5%91%BD%E4%BB%A4%E3%82%BB%E3%83%83%E3%83%88
- 算術演算を実装する
  - ~足し算~
  - ~引き算~
  - ~掛け算~
  - 割り算
- 解説ややったことをまとめる
- ~計算可能性の観点において、チューリングマシンより計算能力のある計算モデルはないんじゃなかったっけ？だからチューリング完全であるBrainf\*ckでは今回作ったCPU作れるし、今回作ったCPUでもBrainf\*ck作れそう~

## 時間があればやりたい
- godoc書きたい
