package store

import "testing"

func TestMemoryCreateListComplete(t *testing.T) {
	m := NewMemory()

	task, err := m.Create("Buy milk")
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if task.ID == "" || task.Title != "Buy milk" || task.Completed {
		t.Fatalf("unexpected task: %+v", task)
	}

	list := m.List()
	if len(list) != 1 || list[0].ID != task.ID {
		t.Fatalf("list: %+v", list)
	}

	done, err := m.Complete(task.ID)
	if err != nil || !done.Completed {
		t.Fatalf("complete: %+v err=%v", done, err)
	}

	again, err := m.Complete(task.ID)
	if err != nil || !again.Completed {
		t.Fatalf("idempotent complete: %+v err=%v", again, err)
	}
}

func TestMemoryCreateEmptyTitle(t *testing.T) {
	m := NewMemory()
	_, err := m.Create("   ")
	if err != ErrTitleRequired {
		t.Fatalf("want ErrTitleRequired, got %v", err)
	}
}

func TestMemoryCompleteNotFound(t *testing.T) {
	m := NewMemory()
	_, err := m.Complete("missing")
	if err != ErrNotFound {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}
