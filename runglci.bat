@echo off
gengo -upd
gengo -env
make revendor
make commit
