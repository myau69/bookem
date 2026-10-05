# Соглашение о названии веток и коммитов

| Тип | Для чего | Пример |
|---|---|---|
| `feat/` | Новая возможность | `feat/room-model` |
| `infra/` | Docker, БД и окружение | `infra/local-postgres` |
| `fix/` | Исправление ошибки | `fix/room-name-validation` |
| `test/` | Отдельная работа над тестами | `test/booking-race` |
| `refactor/` | Перестройка без изменения поведения | `refactor/room-service` |
| `docs/` | Документация | `docs/git-workflow` |
| `chore/` | Обслуживание проекта | `chore/update-dependencies` |
| `ci/` | Автоматические проверки | `ci/project-checks` |

## Коммиты должны быть вида:
```bash
git commit -m "[type]: [description]"
```

$example:$
```bash
git commit -m "feat: add a new feature"
```

## Пуши как коммиты

# Памятка с командами 

```bash
git switch main
git status
git pull --ff-only

git switch -c feat/room-model # работаем с ветками, в мейн не суем все подряд
git branch --show-current

go fmt ./...
go test ./...

git add -- internal/models/rooms.go internal/models/room_test.go
git diff --staged
git commit -m "feat: add room model"
git status

git push -u origin feat/room-model # первый раз пушим так
# ИЛИ
git push # после первого раза пушим так (в одной ветке!!!)

git switch main 
git pull --ff-only
go test ./...
git branch -d feat/room-model # удаляем локальную ветку
git push origin --delete feat/room-model # удаляем ветку на GitHub
git fetch --prune origin # убираем устаревшую ссылку origin/feat/room-model

git switch -c feat/new_branch # создаем новую ветку и повторяем цикл
```
