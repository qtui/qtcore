package qtcore

type QModelIndex struct {
}

type QPersistentModelIndex struct {
}

type QAbstractItemModel struct {
	*QObject
}

func QAbstractItemModelFromptr(ptr voidptr) *QAbstractItemModel {
	return &QAbstractItemModel{QObjectFromptr(ptr)}
}

type QAbstractTableModel struct {
	*QAbstractItemModel
}

func QAbstractTableModelFromptr(ptr voidptr) *QAbstractTableModel {
	return &QAbstractTableModel{QAbstractItemModelFromptr(ptr)}
}

type QAbstractListModel struct {
	*QAbstractItemModel
}

func QAbstractListModelFromptr(ptr voidptr) *QAbstractListModel {
	return &QAbstractListModel{QAbstractItemModelFromptr(ptr)}
}
