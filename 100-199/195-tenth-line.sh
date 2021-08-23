amount=`cat file.txt | wc -l`
if (( amount >= 10 )); then
    cat file.txt | head -n 10 | tail -n 1
fi

# https://leetcode.com/problems/tenth-line/discuss/55544/Share-four-different-solutions
awk 'NR == 10' file.txt
sed -n 10p file.txt
