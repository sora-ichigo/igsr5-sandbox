# Strands SDK Python

Strands SDK for Pythonを使用したAIエージェントのサンプルプロジェクトです。

## セットアップ

### 1. 仮想環境のアクティベート

```bash
source .venv/bin/activate
```

### 2. AWS認証の設定

Strands SDKはAWS Bedrockサービスを使用するため、AWS認証情報の設定が必要です。

#### AWS SSOの設定

```bash
# AWS SSOの設定
aws configure sso
```

設定例：
- SSO session name: `for-testing`
- SSO start URL: `https://d-9567018c51.awsapps.com/start/#`
- SSO region: `ap-northeast-1`

#### ログイン

```bash
# プロファイル名を確認
less ~/.aws/config

# 適切なプロファイルでログイン（例：AdministratorAccess-247679429177）
aws sso login --profile AdministratorAccess-247679429177

# 環境変数の設定
export AWS_PROFILE=AdministratorAccess-247679429177
```

#### 認証確認

```bash
aws sts get-caller-identity
```

### 3. AWS Bedrockモデルアクセス権限の設定

AWS Bedrockで使用するモデルへのアクセス権限が必要です。以下のエラーが出る場合：

```
AccessDeniedException: You don't have access to the model with the specified model ID.
```

**解決方法：**
1. AWS Bedrockコンソールにアクセス
2. 「Model access」セクションで必要なモデルのアクセス権限を有効化
3. 通常、Claude 3 Sonnet、Claude 3 Haikuなどのモデルアクセスを有効化する必要があります

## 使用方法

```bash
python ./agent.py
```

## トラブルシューティング

### よくあるエラー

#### `NoCredentialsError: Unable to locate credentials`
- AWS認証情報が設定されていません
- 上記の「AWS認証の設定」セクションに従って設定してください

#### `AccessDeniedException: You don't have access to the model`
- AWS Bedrockでモデルアクセス権限が設定されていません
- AWS Bedrockコンソールでモデルアクセスを有効化してください

## ファイル構成

- `agent.py`: Strands Agentを使用した数学計算のサンプル
- `.python-version`: 使用するPythonバージョン
- `.venv/`: Python仮想環境
