grep -v -f exclude-from-code-coverage.txt cover.out > temp.out
while read p || [ -n "$p" ]
do
   echo -e "$p:1.1,1.2 1 1" >> temp.out
done < exclude-from-code-coverage.txt
mv temp.out cover.out