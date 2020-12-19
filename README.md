# furiko/gqlgen-todo
このリポジトリは[gqlgen](https://gqlgen.com/getting-started/)自習用の個人リポジトリです。

This repository is created for my own lerning [gqlgen](https://gqlgen.com/getting-started/).

## 起動方法

```bash
make up

make migration/up
```

ブラウザから`localhost:8080`に接続

playgroundを利用してクエリを送信
```graphql
 mutation createTodo {
  createTodo(input:{text:"todo", userId:"1"}) {
    text
    id
    done
  }
}

query findTodos {
    todos {
      text
    	done
      id
    }
}# Write your query or mutation here

mutation deleteTodo {
  deleteTodo(input: {id:"T8674665223082153551"}) { # 存在するidを指定
    text
    id
    done
  }
}

```

## 停止方法
```bash
make down
```
