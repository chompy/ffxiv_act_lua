package main

const assetError404General = "PCFET0NUWVBFIGh0bWw+PGh0bWw+PGhlYWQ+CiAgICAgICAgPHRpdGxlPkZGVEsgLSBOb3QgRm91bmQ8L3RpdGxlPgogICAgICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgdHlwZT0idGV4dC9jc3MiIGhyZWY9ImRhdGE6dGV4dC9jc3M7YmFzZTY0LGFIUnRiQ3dnWW05a2VTQjdDaUFnSUNCbWIyNTBMV1poYldsc2VUb2djMkZ1Y3kxelpYSnBaanNLSUNBZ0lHSmhZMnRuY205MWJtUXRZMjlzYjNJNkl6TTRNREF3TURzS0lDQWdJR052Ykc5eU9pQWpabVptT3dwOUNpNTBhWFJzWlNCN0NpQWdJQ0JzWlhSMFpYSXRjM0JoWTJsdVp6b2dMVFJ3ZURzS0lDQWdJR0p2Y21SbGNpMWliM1IwYjIwNklERndlQ0J6YjJ4cFpDQWpZVFZoTldFMU93b2dJQ0FnZDJsa2RHZzZJRE13TUhCNE93b2dJQ0FnZEdWNGRDMWhiR2xuYmpvZ1kyVnVkR1Z5T3dvZ0lDQWdiV0Z5WjJsdU9pQmhkWFJ2T3dvZ0lDQWdiV0Z5WjJsdUxXSnZkSFJ2YlRvZ016QndlRHNLSUNBZ0lHTnZiRzl5T2lBalptWm1Pd29nSUNBZ1ptOXVkQzF6YVhwbE9pQTJNSEI0T3dvZ0lDQWdabTl1ZEMxbVlXMXBiSGs2SUNkVWFXMWxjeUJPWlhjZ1VtOXRZVzRuTENCVWFXMWxjeXdnYzJWeWFXWTdDaUFnSUNCMGNtRnVjMlp2Y20wNklITmpZV3hsS0M0MUxDQXhLVHNLSUNBZ0lIQnZjMmwwYVc5dU9pQnlaV3hoZEdsMlpUc0tJQ0FnSUhSdmNEb2dObkI0T3dvZ0lDQWdiR1YwZEdWeUxYTndZV05wYm1jNklERndlRHNLQ24wS0xuUnBkR3hsSUhOd1lXNGdld29nSUNBZ1kyOXNiM0k2SUhKbFpEc0tmUXBvTWl3Z2NDQjdDaUFnSUNCMFpYaDBMV0ZzYVdkdU9pQmpaVzUwWlhJN0NpQWdJQ0J0WVhKbmFXNDZJREE3Q2lBZ0lDQnRZWEpuYVc0dFltOTBkRzl0T2lBeE1IQjRPd3A5Q2c9PSIvPgogICAgPC9oZWFkPgogICAgPGJvZHk+CiAgICAgICAgPGgxIGNsYXNzPSJ0aXRsZSI+RkY8c3Bhbj5USzwvc3Bhbj48L2gxPgogICAgICAgIDxoMj5Ob3QgRm91bmQ8L2gyPgogICAgICAgIDxwPlRoZSBwYWdlIHdhcyBub3QgZm91bmQuPC9wPgogICAgCjwvYm9keT48L2h0bWw+"
const assetError404UserNotFound = "PCFET0NUWVBFIGh0bWw+PGh0bWw+PGhlYWQ+CiAgICAgICAgPHRpdGxlPkZGVEsgLSBOb3QgRm91bmQ8L3RpdGxlPgogICAgICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgdHlwZT0idGV4dC9jc3MiIGhyZWY9ImRhdGE6dGV4dC9jc3M7YmFzZTY0LGFIUnRiQ3dnWW05a2VTQjdDaUFnSUNCbWIyNTBMV1poYldsc2VUb2djMkZ1Y3kxelpYSnBaanNLSUNBZ0lHSmhZMnRuY205MWJtUXRZMjlzYjNJNkl6TTRNREF3TURzS0lDQWdJR052Ykc5eU9pQWpabVptT3dwOUNpNTBhWFJzWlNCN0NpQWdJQ0JzWlhSMFpYSXRjM0JoWTJsdVp6b2dMVFJ3ZURzS0lDQWdJR0p2Y21SbGNpMWliM1IwYjIwNklERndlQ0J6YjJ4cFpDQWpZVFZoTldFMU93b2dJQ0FnZDJsa2RHZzZJRE13TUhCNE93b2dJQ0FnZEdWNGRDMWhiR2xuYmpvZ1kyVnVkR1Z5T3dvZ0lDQWdiV0Z5WjJsdU9pQmhkWFJ2T3dvZ0lDQWdiV0Z5WjJsdUxXSnZkSFJ2YlRvZ016QndlRHNLSUNBZ0lHTnZiRzl5T2lBalptWm1Pd29nSUNBZ1ptOXVkQzF6YVhwbE9pQTJNSEI0T3dvZ0lDQWdabTl1ZEMxbVlXMXBiSGs2SUNkVWFXMWxjeUJPWlhjZ1VtOXRZVzRuTENCVWFXMWxjeXdnYzJWeWFXWTdDaUFnSUNCMGNtRnVjMlp2Y20wNklITmpZV3hsS0M0MUxDQXhLVHNLSUNBZ0lIQnZjMmwwYVc5dU9pQnlaV3hoZEdsMlpUc0tJQ0FnSUhSdmNEb2dObkI0T3dvZ0lDQWdiR1YwZEdWeUxYTndZV05wYm1jNklERndlRHNLQ24wS0xuUnBkR3hsSUhOd1lXNGdld29nSUNBZ1kyOXNiM0k2SUhKbFpEc0tmUXBvTWl3Z2NDQjdDaUFnSUNCMFpYaDBMV0ZzYVdkdU9pQmpaVzUwWlhJN0NpQWdJQ0J0WVhKbmFXNDZJREE3Q2lBZ0lDQnRZWEpuYVc0dFltOTBkRzl0T2lBeE1IQjRPd3A5Q2c9PSIvPgogICAgPC9oZWFkPgogICAgPGJvZHk+CiAgICAgICAgPGgxIGNsYXNzPSJ0aXRsZSI+RkY8c3Bhbj5USzwvc3Bhbj48L2gxPgogICAgICAgIDxoMj5Vc2VyIE5vdCBGb3VuZDwvaDI+CiAgICAgICAgPHA+Tm8gdXNlciBmb3VuZCBhdCBnaXZlbiBhZGRyZXNzLCB0aGV5IGNvdWxkIGJlIG9mZmxpbmUuPC9wPgogICAgCjwvYm9keT48L2h0bWw+"
const assetError500General = "PCFET0NUWVBFIGh0bWw+PGh0bWw+PGhlYWQ+CiAgICAgICAgPHRpdGxlPkZGVEsgLSBFcnJvcjwvdGl0bGU+CiAgICAgICAgPGxpbmsgcmVsPSJzdHlsZXNoZWV0IiB0eXBlPSJ0ZXh0L2NzcyIgaHJlZj0iZGF0YTp0ZXh0L2NzcztiYXNlNjQsYUhSdGJDd2dZbTlrZVNCN0NpQWdJQ0JtYjI1MExXWmhiV2xzZVRvZ2MyRnVjeTF6WlhKcFpqc0tJQ0FnSUdKaFkydG5jbTkxYm1RdFkyOXNiM0k2SXpNNE1EQXdNRHNLSUNBZ0lHTnZiRzl5T2lBalptWm1Pd3A5Q2k1MGFYUnNaU0I3Q2lBZ0lDQnNaWFIwWlhJdGMzQmhZMmx1WnpvZ0xUUndlRHNLSUNBZ0lHSnZjbVJsY2kxaWIzUjBiMjA2SURGd2VDQnpiMnhwWkNBallUVmhOV0UxT3dvZ0lDQWdkMmxrZEdnNklETXdNSEI0T3dvZ0lDQWdkR1Y0ZEMxaGJHbG5iam9nWTJWdWRHVnlPd29nSUNBZ2JXRnlaMmx1T2lCaGRYUnZPd29nSUNBZ2JXRnlaMmx1TFdKdmRIUnZiVG9nTXpCd2VEc0tJQ0FnSUdOdmJHOXlPaUFqWm1abU93b2dJQ0FnWm05dWRDMXphWHBsT2lBMk1IQjRPd29nSUNBZ1ptOXVkQzFtWVcxcGJIazZJQ2RVYVcxbGN5Qk9aWGNnVW05dFlXNG5MQ0JVYVcxbGN5d2djMlZ5YVdZN0NpQWdJQ0IwY21GdWMyWnZjbTA2SUhOallXeGxLQzQxTENBeEtUc0tJQ0FnSUhCdmMybDBhVzl1T2lCeVpXeGhkR2wyWlRzS0lDQWdJSFJ2Y0RvZ05uQjRPd29nSUNBZ2JHVjBkR1Z5TFhOd1lXTnBibWM2SURGd2VEc0tDbjBLTG5ScGRHeGxJSE53WVc0Z2V3b2dJQ0FnWTI5c2IzSTZJSEpsWkRzS2ZRcG9NaXdnY0NCN0NpQWdJQ0IwWlhoMExXRnNhV2R1T2lCalpXNTBaWEk3Q2lBZ0lDQnRZWEpuYVc0NklEQTdDaUFnSUNCdFlYSm5hVzR0WW05MGRHOXRPaUF4TUhCNE93cDlDZz09Ii8+CiAgICA8L2hlYWQ+CiAgICA8Ym9keT4KICAgICAgICA8aDEgY2xhc3M9InRpdGxlIj5GRjxzcGFuPlRLPC9zcGFuPjwvaDE+CiAgICAgICAgPGgyPkFuIEVycm9yIEhhcyBPY2N1cmVkPC9oMj4KICAgICAgICA8cD5UaGVyZSB3YXMgYW4gdW5rbm93biBlcnJvciwgcGxlYXNlIHRyeSBhZ2FpbiBsYXRlci48L3A+CiAgICAKPC9ib2R5PjwvaHRtbD4="
const assetFavicon = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAFfUlEQVRYhcWXXWxURRTHf7O73S4LFAqlxU1FwJBQgoAQQNBSqoIKpgZQAw8GEx40EBJ54EEjCcaE8BE/ojES0BeVpAZCAyICBVqxfFOElqCFLpTP3nbZsl8u3bt3Z3zo3rJ79wLliZNs7s75z5nzn3POnLlXTAcfT1BEbXX12+OmTVtiBUKBgBa6c6dlZFlZhRULatrpwMyZjqfhLbtF98JX82CVOW6DmjXwy1b42QNeHYxlsAQwWDnn1XF32ttPSymVlDIhpUxKKVVQ0y6vWbx4WjQUuprG7kkpU2ms5jL8KBHK7vcpvN4MuyRC/Ye4tw7mTgffEfheItR5qJkOvungc5yqPRgKatoBgMa6um8+KH+p/JbfvxvgQHX1zXAw2ABQW129YnVVVWVQ007c35k6sRa1DOAa6tQ61HtJlNQh1gK/AkRQF3fBBYBjsN1AyZ9gg7mGAyCZSEQB8j2eeNPRY9d/WLv2c5lKGQBGMhkD6D9wYKxhz57W+p07e40vwFYPtAHkQ/w3OByAMwAhiACkIGTOL4IBAWhqgOtZBAzDkJk53Ldt2+2gprUCOByOOIBKY1+v+uhsLByOSKAGzlnzH03rUmCkVeaTcTD+NuzNnO9IOzGwSDgYbAVQUjoAhLla0jCCmtbWAfGL0JljB1cBpBUASqBiF+zP1Lls5gEQudvVkUkyC+sKdoQhZefnrg0pgEHglVBcC62PJDClcnZxLByO2GGTyssHR++GIlFw2+EB6LLTr4BZcThi1efsDmDmvPkTbrX6r8P93KefrrlLloy+3tKi/QmH7Wz3wcV/4LZVXwZvHLaEHx4QgYLCwgHfrl59MFMngMJhwzxDSkrcG5cvP2tnB+CHmFUnwKVgbDU0WbGsAgNwOB3uIp+vINOxKeVVVYOdeS7bqD1MPoRpqZ7jmVMzDrgfZoD5S98v7l9Q0JvfzBQ8P7tiRDalvslUeO0qHLfDcnZTuWjRBESuEwGMKhs30c7mYSLBUQxzfTDqkQR0PTFkzKRJS62OAZK6XuobPfrdx92/DmOHIkYUwgQ7PKsIJ1fMXg442tvaaq0TKxYs2GQl3BcZgSgF6P8AAlkLdt640WQkk92A09SZNaBdu3ZMKWVk1ktfJJV+Dkb4RkKRFc+KQOfNm7ujodCJDLte8Tc3f0fGxdJXaUedKoYJXoTnHRi/Ceoz8ZxjeP5ow++ZE0xMCCFv+v1/CPs2/0BxQygKlwCehYlWPOcY1mzefEVJmeNEAZfPnWt93BQARNMNaIhNHeQUVWtTc6Srs7M70zH0RKKxrq4T9fgUOuA82BeiXVXLcCDQexFlpqexri6W1BOPkYIesqfTERiE8I2xFKLtsYpFIt12+u54XL/b2an3nUAP/e1wKQ7dTmChJQq2BE7u33/B/G+5DY3DO3a09Z1Aj0jQo6h/AUZbCtGWwOlDhzSrzkxFY1297X3/KDEL0doRH9nZhBAPzXlfW3MHNAMMsCMgUyk3gCsvL2lDAACllO0bkNkylWUzDsu4GS4CDEIMn5zxNeYAyMvP7xk4nXlWB0opbxrz2BHol341E5au6r4/dgPsgrZU2uFCeCGLgMfrfSr9LLY6cOfnFwEMLCzM6eMApVAM4IThmfoSGA4KVxoPQVxHSYBRMMec51y/cePUKZWVn7jz8wvcHk/poKKiUycPHNAAPt6yZdZzM2asdLpcHo/XW+oQov780aNh0/hFGL4QNhQgSvohCp+Be/Vw5mXwvQnrCxDDPIihJXDjFZhYhpgL4EWMcsJff0O7K5lI6C1nG79QqqegDF3vPed6IhG7cPz4Z+ZRTCb1rO+HFMgrcPAK6iBAAgKm3g/7/Kh9AHHoioJxBPWladsNcQDxpD/P/wdIYjDmz0IZTwAAAABJRU5ErkJggg=="
const assetIndex = "PGh0bWw+PGhlYWQ+PC9oZWFkPjxib2R5PjwvYm9keT48L2h0bWw+"
