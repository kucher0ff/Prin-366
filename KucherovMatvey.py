# Простой менеджер задач (без сохранения в файл)

tasks = []  # Каждая задача: {'description': str, 'completed': bool}

def show_tasks():
    """Вывести все задачи с номерами и статусом."""
    if not tasks:
        print("\nСписок задач пуст.")
        return
    print("\nТекущие задачи:")
    for i, task in enumerate(tasks, 1):
        status = '✓' if task['completed'] else '✗'
        print(f"{i}. [{status}] {task['description']}")

def add_task():
    """Добавляет новую задачу."""
    desc = input("Введите описание задачи: ").strip()
    if desc:
        tasks.append({'description': desc, 'completed': False})
        print("Задача добавлена.")
    else:
        print("Описание не может быть пустым.")

def complete_task():
    """Отмечает задачу как выполненную."""
    show_tasks()
    if not tasks:
        return
    try:
        num = int(input("Введите номер задачи для отметки: "))
        if 1 <= num <= len(tasks):
            tasks[num-1]['completed'] = True
            print("Задача отмечена выполненной.")
        else:
            print("Неверный номер.")
    except ValueError:
        print("Ошибка: введите целое число.")

def delete_task():
    """Удаляет задачу по номеру."""
    show_tasks()
    if not tasks:
        return
    try:
        num = int(input("Введите номер задачи для удаления: "))
        if 1 <= num <= len(tasks):
            deleted = tasks.pop(num-1)
            print(f"Задача '{deleted['description']}' удалена.")
        else:
            print("Неверный номер.")
    except ValueError:
        print("Ошибка: введите целое число.")

def main():
    while True:
        print("\n--- МЕНЕДЖЕР ЗАДАЧ ---")
        print("1. Показать все задачи")
        print("2. Добавить задачу")
        print("3. Отметить как выполненную")
        print("4. Удалить задачу")
        print("5. Выход")
        choice = input("Выберите действие: ").strip()

        if choice == '1':
            show_tasks()
        elif choice == '2':
            add_task()
        elif choice == '3':
            complete_task()
        elif choice == '4':
            delete_task()
        elif choice == '5':
            print("До свидания!")
            break
        else:
            print("Неверный ввод. Пожалуйста, выберите пункт из меню.")

if __name__ == "__main__":
    main()