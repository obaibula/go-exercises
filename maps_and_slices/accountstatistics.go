package maps_and_slices

// PartitionAccountsBySex partitions a slice of Accounts into a map based on their sex.
// The resulting map contains two keys: true for Male accounts and false for Female accounts.
//
// Params:
//   - accounts: A slice of Account structures to partition by sex.
//
// Returns:
//   - A map where the keys are booleans indicating the sex (true for Male, false for Female),
//     and the values are slices of Account structures corresponding to each sex.
func PartitionAccountsBySex(accounts []Account) map[bool][]Account {
	panic("not implemented")
}

// FindRichestPerson finds and returns the account with the highest balance from a slice of Accounts.
// If the slice is empty, it returns false using the comma-ok idiom.
//
// In case multiple accounts have the same highest balance, the function uses a provided mergeFunction
// to resolve the tie. If no mergeFunction is provided (i.e., it is nil), the default behavior is to
// resolve the tie by selecting the account with the lower ID, using prepared AccountIdBinaryOperator.
//
// Params:
//   - accounts: A slice of Account structures to search through.
//   - mergeFunction: A custom function (AccountBinaryOperator) that determines how to resolve ties
//     between accounts with equal balances. If nil, AccountIdBinaryOperator is used.
//
// Returns:
// - The Account with the highest balance.
// - A boolean indicating whether the slice contained any accounts (false if the slice was empty).
func FindRichestPerson(accounts []Account, mergeFunction AccountBinaryOperator) (Account, bool) {
	panic("not implemented")
}

type AccountBinaryOperator func(left, right Account) Account

func AccountIdBinaryOperator(left, right Account) Account {
	if left.ID < right.ID {
		return left
	} else {
		return right
	}
}
