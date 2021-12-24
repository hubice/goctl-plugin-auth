root_path=./
for dir in  $root_path/*; do
  if test -f $dir;  then
    echo "-"
  else
    for file in $dir/*; do
      admin=$(echo $file | grep "admin.api")
      if [[ $admin != "" ]] ; then
        goctl api plugin -p goctl-auth-api="goctl-auth-api" -api $file -dir .
      fi
    done
  fi
done
sleep 5