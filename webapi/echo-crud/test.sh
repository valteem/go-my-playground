# -X specifies request method

for i in `seq 1 10`
do 
    curl -X POST -H 'Content-Type: application/json' -d '{"desc":"someNewItem"}' 127.0.0.1:44567/items
done

for i in `seq 11 15`
do 
    curl -X GET -H 'Content-Type: application/json' "127.0.0.1:44567/items/${i}"
done

for i in `seq 6 10`
do 
    curl -X PUT -H 'Content-Type: application/json' -d '{"desc":"newItemDescription"}' "127.0.0.1:44567/items/${i}"
done

for i in `seq 8 11`
do 
    curl -X DELETE -H 'Content-Type: application/json' "127.0.0.1:44567/items/${i}"
done