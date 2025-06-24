#pragma once

#include "binary.h"

#include <bpf/types.h>

enum unwind_type : u8 {
    UNWIND_TYPE_DISABLED = 0,
    UNWIND_TYPE_FP = 1,
    UNWIND_TYPE_DWARF = 2,
};

enum interpreter_type : u8 {
    INTERPRETER_TYPE_NONE = 0,
    INTERPRETER_TYPE_PYTHON = 1,
};

struct interpreter_binary {
    binary_id id;
    enum interpreter_type type;
};

struct process_info {
    enum unwind_type unwind_type;
    binary_id main_binary_id;

    // Interpreter binary ID may be different from main binary ID.
    // For example, if CPython is dynamically linked
    // into the main binary.
    struct interpreter_binary interpreter_binary;
    u64 interpreter_binary_start_address;
};
