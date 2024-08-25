package product_status

// TODO здесь необходимо использовать enum.Enum
// Константы лучше оставить в том виде, в котором они есть. Но их надо занести в допустимые значения Enum.
// Также в этом пакете должен быть доступ к получению используемого Enum.
// Необходимо реализовать тип Status, который будет знать своё значение (а в идеале ещё и свой ReadableString).
// Ну и не лишней будет возможность получать из строки тип Status.

// товар на этапе формирования (черновик)
const Draft = "draft"

// товар на этапе модерации
const Moderation = "moderation"

// товар активен и доступен для покупки
const Active = "active"

// товар не в наличии
const OutOfStock = "out_of_stock"

func ReadableString(status string) string {
	// FIXME неудобно добавлять новые значения статуса и их можно забыть добавить в эту функцию
	if status == Draft {
		return "Draft"
	}
	if status == Moderation {
		return "On moderation"
	}
	if status == Active {
		return "Active"
	}
	if status == OutOfStock {
		return "Out of stock"
	}
	return "Unknown"
}

func Exists(status string) bool {
	// FIXME при добавлении нового значения статуса можно забыть добавить его в эту функцию
	return status == Draft || status == Moderation || status == Active || status == OutOfStock
}
