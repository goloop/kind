# kind - довідник

`kind` - кешований помічник рефлексії для класифікації Go-значень і статичних
типів. Орієнтований на Go 1.24+, без сторонніх залежностей.

Англійська версія: **[DOC.md](DOC.md)**.

## Мета

`kind` відповідає на одне питання - **«що це за тип і що він уміє?»** - для
коду, який мусить обробляти *довільні* Go-типи під час виконання. Він загортає
`reflect` у кешований іммутабельний дескриптор із пласким словником предикатів,
тож питати дешево, а помилитися важко.

## Коли доцільно використовувати

Модуль виправдовує себе тоді, коли твій код отримує типи, яких не може знати на
етапі компіляції, і мусить ухвалювати щодо них рішення:

- **Парсери, декодери й біндери** - головна аудиторія. Змінні оточення, файли
  конфігурації, CLI-прапорці, query-параметри, рядки БД: усе, що бере текст і
  вкладає його в структуру, яку визначає *викликач*. Ця робота зазвичай
  виливається в купу ручного `reflect` («це int? вказівник на структуру? який
  тип елемента за цими двома слайсами?»); із `kind` диспетчеризація читається
  як таблиця:

  ```go
  k := kind.Of(target)
  switch {
  case k.CanImplement(textUnmarshalerType): // рахується й *T, не лише T
      // передати сирий текст в UnmarshalText
  case k.IsAnyInt():
      n, _ := strconv.ParseInt(raw, 10, 64)
      // виставити як int/int8/.../int64
  case k.IsSlice():
      // розбити й рекурсувати по k.Elem()
  }
  ```

- **Виявлення здатностей.** «Чи реалізує цей тип - або вказівник на нього -
  `encoding.TextUnmarshaler`, `sql.Scanner`, `flag.Value`, метод
  `Set(string) error`?» `CanImplement` кодує правило pointer-receiver, яке
  парсерам реально потрібне, а `Is*`-предикати здатностей покривають поширені
  stdlib-інтерфейси одним стилем.
- **Обхід структур.** Списки полів, теги й вкладені форми для валідаторів,
  генераторів схем і коду.
- **Гарячі шляхи.** Дескриптори кешуються за `reflect.Type`, тож цикл парсингу
  платить за класифікацію один раз на тип, а не на кожне значення.

## Коли НЕ використовувати

- **Типи відомі на етапі компіляції.** Звичайний type switch кращий за
  будь-яку рефлексію - зокрема й за `kind`.
- **Одна дрібна перевірка в одному місці.** Два рядки stdlib `reflect` не
  виправдовують залежності.
- **Треба записувати значення.** `kind` - інтроспекція лише на читання: він
  каже, що це за тип і що той уміє, але запис значення лишається твоїм кодом
  (`reflect.Value.Set`, згенеровані сетери тощо).
- **(Де)серіалізація стандартного формату.** `encoding/json` і компанія вже
  роблять власну рефлексію; `kind` там нічого не додає.

Коротко: якщо ти ніколи не тягнешся по `reflect`, тобі `kind` не потрібен. Коли
ж тягнешся - щоб приймати довільні типи користувача в парсері чи біндері - він
прибирає нудну, схильну до помилок половину роботи.

## Конструктори

```go
func Of(v any) *Kind
func OfType(t reflect.Type) *Kind
func TypeOf[T any]() *Kind
```

- `Of` інспектує динамічний тип і стан значення `v`.
- `OfType` інспектує `reflect.Type` без значення.
- `TypeOf[T]` інспектує статичний генерик-тип і є найкращим способом питати про
  інтерфейсні типи.

## Основні методи

```go
func (k *Kind) Type() reflect.Type
func (k *Kind) Kind() reflect.Kind
func (k *Kind) Name() string
func (k *Kind) String() string
func (k *Kind) Value() any
func (k *Kind) Is(name string) bool
func (k *Kind) Elem() *Kind
func (k *Kind) Key() *Kind
func (k *Kind) MapKeyKind() *Kind
func (k *Kind) MapValueKind() *Kind
func (k *Kind) Leaf() *Kind
func (k *Kind) Base() *Kind
```

`Elem` працює для вказівників, масивів, слайсів, map і каналів. `Key` працює для
map. Якщо такого компонента типу немає, методи повертають nil-Kind, чий `Name()`
дорівнює `"nil"`.

`Is` ігнорує пробіли в порівнюваному імені. Він кейс-нечутливий лише для
вбудованих скалярних імен на кшталт `int`, `string`, `float64`; іменовані й
складені типи порівнюються з урахуванням регістру.

## Предикати

Стан значення:

```go
IsNil() bool
IsZero() bool
IsNilable() bool
IsNamed() bool
IsEmpty() bool
IsTruthy() bool
```

`IsEmpty` повертає true для nil, нульового значення й рядків/масивів/слайсів/map/
каналів довжини нуль. Він **не** розіменовує вказівники: ненульовий вказівник на
нульове значення не порожній (`Of(&T{}).IsEmpty() == false`). Виклич спершу
`Deref`, якщо потрібна порожність того, на що вказує вказівник.

Контейнери та складені/референсні типи:

```go
IsScalar() bool
IsContainer() bool
IsComposite() bool
IsReference() bool
IsComparable() bool
IsOrdered() bool
IsNumeric() bool
IsPointer() bool
IsPointerToStruct() bool
IsArray() bool
IsSlice() bool
IsSliceLike() bool
IsSliceOfSlices() bool
IsArrayOfSlices() bool
IsSliceOfArrays() bool
IsArrayOfArrays() bool
IsMap() bool
IsMapLike() bool
IsStruct() bool
IsInterface() bool
IsFunction() bool
IsChannel() bool
IsUnsafePointer() bool
IsComplex() bool
```

Скаляри:

```go
IsBool() bool
IsString() bool
IsInt() bool
IsInt8() bool
IsInt16() bool
IsInt32() bool
IsInt64() bool
IsUint() bool
IsUint8() bool
IsUint16() bool
IsUint32() bool
IsUint64() bool
IsUintptr() bool
IsFloat32() bool
IsFloat64() bool
IsComplex64() bool
IsComplex128() bool
IsNumber() bool
IsAnyInt() bool
IsAnyFloat() bool
IsAnyComplex() bool
IsUnsigned() bool
IsSigned() bool
```

Скалярні предикати **leaf-aware**: контейнер повідомляє тип листка свого ланцюга
елементів, тож `kind.Of([]int{}).IsSlice()` і `kind.Of([]int{}).IsInt()` обидва
повертають true. Для map скалярні leaf-предикати описують лише тип значення;
ключ інспектуй через `Key` або `MapKeyKind`.

Саме цього хоче рекурсивний парсер (зазирнути в листок, тоді спускатися). Коли ж
потрібна строга перевірка, що значення - *саме* скаляр, а не контейнер,
порівнюй `Kind`:

```go
k.Kind() == reflect.Int // строго int, а не []int чи map[K]int
```

## Форма й непрямість

```go
func (k *Kind) ChanDir() reflect.ChanDir
func (k *Kind) Len() int
func (k *Kind) Cap() int
func (k *Kind) Depth() int
func (k *Kind) Deref() *Kind
func (k *Kind) Indirect() *Kind
func (k *Kind) PointerDepth() int
func (k *Kind) IsPointerTo(target reflect.Type) bool
func IsPointerTo[T any](k *Kind) bool
```

`Leaf` / `Base` йдуть за елементами вказівника, масиву, слайса, map і каналу до
термінального типу. Наприклад `[][]*User` має глибину `3` і листок `User`.

`Len` повертає довжину значення для рядків, масивів, слайсів, map і каналів. Для
type-only масивів повертає довжину масиву. Інакше повертає `-1`. `Cap` повертає
місткість для масивів, слайсів і каналів, або `-1`.

## Поля структур

```go
type Field struct {
    Name      string
    Type      *Kind
    Index     []int
    Tag       reflect.StructTag
    Anonymous bool
    Exported  bool
    Offset    uintptr
}

func (k *Kind) Fields() []Field
func (k *Kind) ExportedFields() []Field
func (k *Kind) HasField(name string) bool
func (k *Kind) Field(name string) (Field, bool)
func (k *Kind) HasTag(key string) bool
func (k *Kind) FieldsByTag(key string) []Field
func (f Field) HasTag(key string) bool
func (f Field) TagValue(key string) (string, bool)
```

API полів інспектують безпосередні поля структури. Повернені слайси й значення
індексів - захисні копії, тож викликач може сортувати чи змінювати їх, не
чіпаючи кеш.

## Присвоюваність та інтерфейси

Go не підтримує генерик-методи, тож генерик-хелпери - це top-level функції, а не
методи `k.Implements[T]()`.

```go
func (k *Kind) Implements(target reflect.Type) bool
func (k *Kind) CanImplement(target reflect.Type) bool
func (k *Kind) AssignableTo(target reflect.Type) bool
func (k *Kind) ConvertibleTo(target reflect.Type) bool
func Implements[T any](k *Kind) bool
func CanImplement[T any](k *Kind) bool
func AssignableTo[T any](k *Kind) bool
func ConvertibleTo[T any](k *Kind) bool
```

`Implements` строгий: інтерфейс має реалізовувати саме цей тип. `CanImplement`
орієнтований на здатність: інтерфейс може реалізовувати цей тип або, для
не-вказівникових типів, `*T`. Хелпери здатності беруть це практичне правило, бо
парсери й unmarshaler-методи часто мають pointer-receiver.

Для поширених перевірок парсера/кодування:

```go
IsTextMarshaler() bool
IsTextUnmarshaler() bool
IsBinaryMarshaler() bool
IsBinaryUnmarshaler() bool
IsJSONMarshaler() bool
IsJSONUnmarshaler() bool
IsError() bool
IsStringer() bool
IsEnvMarshaler() bool
IsEnvUnmarshaler() bool
IsValidator() bool
IsVerifier() bool
IsStringParser() bool
IsBytesParser() bool
IsSetter() bool
IsFlagValue() bool
IsScanner() bool
IsValuer() bool
IsReader() bool
IsWriter() bool
IsCloser() bool
IsReaderFrom() bool
IsWriterTo() bool
IsLogValuer() bool
```

Capability-інтерфейси в GoLoop-стилі - це локальні форми, а не імпорти з інших
GoLoop-модулів:

```go
type EnvMarshaler interface {
    MarshalEnv() (map[string]string, error)
}

type EnvUnmarshaler interface {
    UnmarshalEnv(map[string]string) error
}

type Validator interface {
    Validate() error
}

type Verifier interface {
    Valid() bool
}

type StringParser interface {
    Parse(string) error
}

type BytesParser interface {
    ParseBytes([]byte) error
}

type Setter interface {
    Set(string) error
}
```

## Конверсії

```go
AsBool() (bool, bool)
AsString() (string, bool)
AsInt() (int, bool)
AsInt8() (int8, bool)
AsInt16() (int16, bool)
AsInt32() (int32, bool)
AsInt64() (int64, bool)
AsUint() (uint, bool)
AsUint8() (uint8, bool)
AsUint16() (uint16, bool)
AsUint32() (uint32, bool)
AsUint64() (uint64, bool)
AsUintptr() (uintptr, bool)
AsFloat32() (float32, bool)
AsFloat64() (float64, bool)
AsComplex64() (complex64, bool)
AsComplex128() (complex128, bool)
AsUnsafePointer() (unsafe.Pointer, bool)
```

Методи `As*` повертають `ok=false`, коли динамічне значення не є запитаним типом.
Іменовані скалярні типи конвертуються у свою вбудовану форму замість паніки.

## Приклади

```go
k := kind.Of(map[string][]int{"one": {1, 2, 3}})

fmt.Println(k.IsMap())                  // true
fmt.Println(k.MapKeyKind().IsString())  // true
fmt.Println(k.MapValueKind().IsSlice()) // true
fmt.Println(k.MapValueKind().IsInt())   // true
fmt.Println(k.IsString())               // false
```

```go
type Reader interface {
    Read([]byte) (int, error)
}

k := kind.TypeOf[Reader]()
fmt.Println(k.IsInterface()) // true
```

```go
type Config struct {
    Port int `env:"PORT" json:"port"`
}

k := kind.TypeOf[Config]()
field, _ := k.Field("Port")
fmt.Println(k.HasTag("env"))          // true
fmt.Println(field.Type.IsInt())       // true
fmt.Println(field.TagValue("json"))   // "port", true
```

```go
type Request struct{}

func (*Request) Validate() error { return nil }

k := kind.TypeOf[Request]()
fmt.Println(k.IsValidator()) // true: *Request має Validate
```

## Продуктивність

Пакет кешує дескриптори за `reflect.Type` у `sync.Map`. `Of(value)` усе одно
обчислює стан значення (`nil`, `zero`) на кожен виклик. `OfType`, `TypeOf`,
`Elem`, `Key`, `MapKeyKind` і `MapValueKind` перевикористовують type-only Kind з
кешу.
