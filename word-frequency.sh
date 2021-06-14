cat words.txt | tr -s "[:blank:]" "\n" | sed '/^$/d' | sort | uniq -c | sort -nr | awk '{print $2,$1}'
# tr -s 将连续出现的多个匹配视为一个，这样可以省去 sed 删空行的步骤，运行时间从 4ms 变为 0ms
cat words.txt | tr -s "[:blank:]" "\n" | sort | uniq -c | sort -nr | awk '{print $2,$1}'

# 找了个用 awk 统计词频的题解
awk '{for (i = 1; i <= NF; i++) {cnt[$i]++}} END {for (w in cnt) {print w,cnt[w]}}' words.txt | sort -k2nr
