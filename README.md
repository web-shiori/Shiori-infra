# Shiori-infra
Web SnapshotのAWSリソースのテンプレートを管理するリポジトリ

## デプロイ
```
aws cloudformation update-stack --stack-name (スタック名) --template-body file://web-shiori.yaml --capabilities CAPABILITY_NAMED_IAM
```
