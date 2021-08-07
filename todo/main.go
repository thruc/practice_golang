package main

import (
	"fmt"

	"todo/app/controllers"
	"todo/app/models"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}

/*
次のレクチャーでの、パッケージをimportして実行した際に

is not in GOROOT

表題のエラーが出てしまい、他のフォルダから変数の参照ができない場合。



ターミナルで実行してみてください。



go mod init



go mod tidy





下記記事が参考になるかと思います。



https://qiita.com/taku-yamamoto22/items/4d6f9ff8451a0b86997b





go 1.13からはgo modでモジュール管理が推奨されていました。



プロジェクト内で

go mod init で、依存モジュールを初期化しますと

go.modというファイルが作成されます。中身は、以下のような感じです。

module example.com/go-mod-test go 1.16
依存モジュールの情報は go.mod と go.sum という名前のファイルに記載されます。

これらのファイルを git などでバージョン管理することによって、依存モジュールとそのバージョンを明確にすることができます。



go buildなどのビルドコマンドで、依存モジュールを自動インストールする

go list -m all で、現在の依存モジュールを表示する

go get で、依存モジュールの追加やバージョンアップを行う

go mod tidy で、使われていない依存モジュールを削除する

実は go mod を直接実行することは少なく、他の go サブコマンドを実行したときに、自動的に処理が行われることが多いです。



例えば、後のレクチャーで出てくるiniパッケージを使う場合は、

プロジェクトのルートで

go get -u "gopkg.in/go-ini/ini.v1"
を実行しますと、go.modにも依存パッケージとして記載され使用することができるかと思います。





公式の説明は以下になります

GO111MODULE = onの場合、goコマンドではモジュールを使用する必要があり、GOPATHを参照することはありません。これを、モジュール対応または「モジュール対応モード」で実行されるコマンドと呼びます。

GO111MODULE = offの場合、goコマンドはモジュールサポートを使用しません。代わりに、ベンダーのディレクトリとGOPATHを調べて依存関係を見つけます。これを「GOPATHモード」と呼びます。

GO111MODULE = autoまたは未設定の場合、goコマンドは現在のディレクトリに基づいてモジュールサポートを有効または無効にします。モジュールのサポートは、現在のディレクトリにgo.modファイルが含まれる場合、またはgo.modファイルが含まれるディレクトリの下にある場合にのみ有効になります。



モジュール対応モードでは、GOPATHはビルド中のインポートの意味を定義しなくなりましたが、ダウンロードされた依存関係（GOPATH / pkg / mod）とインストールされたコマンド（GOBINが設定されていない場合はGOPATH / bin）を保存します。


*/
