# rpc article

## 运维

```shell
ops --variable .cfg/dev.yaml -a run --env dev --task codegen
ops --variable .cfg/dev.yaml -a run --env dev --task image
ops --variable .cfg/dev.yaml --env dev -a run --task helm --cmd=diff
ops --variable .cfg/dev.yaml --env dev -a run --task helm --cmd=install
ops --variable .cfg/dev.yaml --env dev -a run --task helm --cmd=upgrade
```