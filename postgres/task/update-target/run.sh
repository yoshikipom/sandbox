while read id
do
if [ -n "${id}" ]; then
    psql -h localhost -p 5432 -U root person -v target_id=${id} -f update.sql
fi
done < id_list
