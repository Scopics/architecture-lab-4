# architecture-lab-4

## Пакет engine

#### commandQueue

- Черга включає в себе `sync.Mutex` для забезпечення контролю доступу до спільного слайсу команд з різних рутин, він нам буде гарантувати, що між `Lock()` та `Unlock()` тільки одна рутина матиме доступ до ресурсу.

- Метод `pull()` реалізовано таким чином, що він спочатку викликає `Lock()`, потім перевіряє чи пустий слайс з командами, якщо це так, то встановлює `noCommand` в `true`, та викликає `Unlock()`, щоб далі чекати значення з каналу `pushed` у який, у разі додавання нової команди приходить сигнал. Як тільки ми дочекались знову викликається Lock() та береться елемент з слайсу. Така імплементація блокує цикл виконання `EventLoop` поки черга пуста, але вже хтось запустив цикл. В кінці завдяки defer, виконається `Unlock()`

- Метод `push()` також викликає `Lock()`, потім додає елемент в слайс і перевіряє, чи чекає хтось на новий елемент в слайсі, якщо це так, то відравляє у канал `pushed` сигнал.
В кінці завдяки defer, виконається `Unlock()`

- Метод `length()` просто повертає довжину слайса команд, але робить це безпечно завдяки використання `Mutex`. Він знадобиться при перевірці умови виходу з циклу подій.

#### cmdExecutor

- Це допоміжна внутрішня команда, що реазіує інтерфейс `Command`, яка знадобиться нам, щоб безпечно вставновлювати значення змінної про фініш циклу, адже ми читаємо прапор фінішу з іншої рутини, тоді як виставляти можемо з нашої основної. В даній стурктурі ми передаємо при ініціалізації функцію, котра викличеться в методі `Execute()`

#### EventLoop

- EventLoop займається асинхронним виконанням команд, які можна додати через метод `Post()`, який просто викличе `push()` на черзі команд

- При виклику `Start()` ми ініціалзізуємо чергу, та канал в який ми будемо посилати сигнал у разі коли цикл подій опрацював всі події та хтось чекає на завершення. Далі в іншій рутині запусається цикл котрий працює доти, доки черга не пуста та ніхто не чекає на завершення, в тілі циклу ми отримуємо команду за допомогою `pull()` (як говорилось вище він заблокує потік виконання, якщо команди немає), і дочекавшись з нього команду викликаємо на ній метод `Execute()`, якщо ми вийшли з циклу нам потрібно ще надіслати значення в канал, що ми закінчили і той хто нас чекає міг теж завершитись

- І нарешті метод `AwaitFinish()` нам потрібен для того, щоб сказати циклу подій, що хтось чекає на те, щоб цикл завершивсь, але прапорець про завершення просто так не можна поставити в `true`, тому ми й користуємось спеціальною командою, яка описана вище. Опублікувавши команду, що встановить прапорець в значення, що хтось чекає фінішу, ми очікуємо значення з каналу, який сповістить нам що ми завершились та виходимо. 

## Пакет commands

- Метод `parse()` розбиває вхідний рядок на команду та її параметри, і викликає з ними `compose()`.

- Метод `compose()` знаходить зі списку структуру команди, відповідно до аргументу-імені. Потім копіює та встановлює поля структури за допомогою `setArgs()`.

- Метод `setField()` встановлює окреме поле структури команди.
