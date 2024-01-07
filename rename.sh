
if [ $# -eq 0 ]
then
  echo "Usage: $0 name_of_go_package"
  exit 1
fi

INITIAL_NAME="github.com/w1png/telegram-bot-template"

find . -name "*.go" -exec sed -i '' -e "s/$INITIAL_NAME/$1/g" {} \;
sed -i '' -e "s/$INITIAL_NAME/$1/g" go.mod
