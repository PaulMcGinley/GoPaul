package array

import "testing"

func TestAppend(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = Append(arr, 4, 5, 6)
	if len(arr) != 6 {
		t.Errorf("Expected length of 6, got %d", len(arr))
	}
	if arr[3] != 4 {
		t.Errorf("Expected 4, got %d", arr[3])
	}
	if arr[4] != 5 {
		t.Errorf("Expected 5, got %d", arr[4])
	}
	if arr[5] != 6 {
		t.Errorf("Expected 6, got %d", arr[5])
	}
}

func TestPrepend(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = Prepend(arr, 4, 5, 6)
	if len(arr) != 6 {
		t.Errorf("Expected length of 6, got %d", len(arr))
	}
	if arr[0] != 4 {
		t.Errorf("Expected 4, got %d", arr[0])
	}
	if arr[1] != 5 {
		t.Errorf("Expected 5, got %d", arr[1])
	}
	if arr[2] != 6 {
		t.Errorf("Expected 6, got %d", arr[2])
	}
}

func TestInsert(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = Insert(arr, 1, 4, 5, 6)
	if len(arr) != 6 {
		t.Errorf("Expected length of 6, got %d", len(arr))
	}
	if arr[1] != 4 {
		t.Errorf("Expected 4, got %d", arr[1])
	}
	if arr[2] != 5 {
		t.Errorf("Expected 5, got %d", arr[2])
	}
	if arr[3] != 6 {
		t.Errorf("Expected 6, got %d", arr[3])
	}
}

func TestDeleteElement(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = DeleteElement(arr, 1)
	if len(arr) != 2 {
		t.Errorf("Expected length of 2, got %d", len(arr))
	}
	if arr[0] != 1 {
		t.Errorf("Expected 1, got %d", arr[0])
	}
	if arr[1] != 3 {
		t.Errorf("Expected 3, got %d", arr[1])
	}
}

func TestCopy(t *testing.T) {
	arr := []int{1, 2, 3}
	result := Copy(arr)
	if len(result) != 3 {
		t.Errorf("Expected length of 3, got %d", len(result))
	}
	if result[0] != 1 {
		t.Errorf("Expected 1, got %d", result[0])
	}
	if result[1] != 2 {
		t.Errorf("Expected 2, got %d", result[1])
	}
	if result[2] != 3 {
		t.Errorf("Expected 3, got %d", result[2])
	}
}

func TestCopyEmpty(t *testing.T) {
	arr := []int{}
	result := Copy(arr)
	if len(result) != 0 {
		t.Errorf("Expected length of 0, got %d", len(result))
	}
}

func TestCopyString(t *testing.T) {
	arr := []string{"a", "b", "c"}
	result := Copy(arr)
	if len(result) != 3 {
		t.Errorf("Expected length of 3, got %d", len(result))
	}
	if result[0] != "a" {
		t.Errorf("Expected a, got %s", result[0])
	}
	if result[1] != "b" {
		t.Errorf("Expected b, got %s", result[1])
	}
	if result[2] != "c" {
		t.Errorf("Expected c, got %s", result[2])
	}
}

func TestCopyStringEmpty(t *testing.T) {
	arr := []string{}
	result := Copy(arr)
	if len(result) != 0 {
		t.Errorf("Expected length of 0, got %d", len(result))
	}
}

func TestCopyFloat(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3}
	result := Copy(arr)
	if len(result) != 3 {
		t.Errorf("Expected length of 3, got %d", len(result))
	}
	if result[0] != 1.1 {
		t.Errorf("Expected 1.1, got %f", result[0])
	}
	if result[1] != 2.2 {
		t.Errorf("Expected 2.2, got %f", result[1])
	}
	if result[2] != 3.3 {
		t.Errorf("Expected 3.3, got %f", result[2])
	}
}

func TestRemove(t *testing.T) {
	arr := []int{1, 2, 3}
	arr = Remove(arr, 2)
	if len(arr) != 2 {
		t.Errorf("Expected length of 2, got %d", len(arr))
	}
	if arr[0] != 1 {
		t.Errorf("Expected 1, got %d", arr[0])
	}
	if arr[1] != 3 {
		t.Errorf("Expected 3, got %d", arr[1])
	}
}

func TestRemoveEmpty(t *testing.T) {
	arr := []int{}
	arr = Remove(arr, 0)
	if len(arr) != 0 {
		t.Errorf("Expected length of 0, got %d", len(arr))
	}
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3}
	if !Contains(arr, 2) {
		t.Errorf("Expected true, got false")
	}
}

func TestContainsEmpty(t *testing.T) {
	arr := []int{}
	if Contains(arr, 0) {
		t.Errorf("Expected false, got true")
	}
}

func TestGetElementIndex(t *testing.T) {
	arr := []int{1, 2, 3}
	index := GetElementIndex(arr, 2)
	if index != 1 {
		t.Errorf("Expected 1, got %d", index)
	}
}

func TestGetElementIndexEmpty(t *testing.T) {
	arr := []int{}
	index := GetElementIndex(arr, 0)
	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}
}

func TestGetElementIndexNotFound(t *testing.T) {
	arr := []int{1, 2, 3}
	index := GetElementIndex(arr, 4)
	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}
}

func TestDistinct(t *testing.T) {
	arr := []int{1, 2, 3, 2, 1}
	result := Distinct(arr)
	if len(result) != 3 {
		t.Errorf("Expected length of 3, got %d", len(result))
	}
	if result[0] != 1 {
		t.Errorf("Expected 1, got %d", result[0])
	}
	if result[1] != 2 {
		t.Errorf("Expected 2, got %d", result[1])
	}
	if result[2] != 3 {
		t.Errorf("Expected 3, got %d", result[2])
	}
}

func TestDistinctEmpty(t *testing.T) {
	arr := []int{}
	result := Distinct(arr)
	if len(result) != 0 {
		t.Errorf("Expected length of 0, got %d", len(result))
	}
}
