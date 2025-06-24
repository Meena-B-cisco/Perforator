#pragma once

#include "binary.h"

struct python_interpreter_state_offsets {
    u32 next_offset;
    u32 threads_head_offset;
};

struct python_runtime_state_offsets {
    u32 py_interpreters_main_offset;
};

struct python_thread_state_offsets {
    u32 cframe_offset;
    u32 current_frame_offset;
    u32 native_thread_id_offset;
    u32 prev_thread_offset;
    u32 next_thread_offset;
};

// Python threads share the same thread group in current pid namespace.
// We also use innermost namespace pid because it is native_thread_id field on PyThreadState structure.
// This allows caching *PyThreadState for docker processes which have host thread group id
// as first part of a key and container pid as the second part of a key.
// First part of a key is needed to avoid collisions with other processes.
// For example collision can happen if we have 2 python containers where both python processes have pid 1.
struct python_thread_key {
    u32 current_ns_pid;
    u32 inner_ns_tid;
};

enum {
    MAX_PYTHON_THREADS = 16384,
    MAX_PYTHON_THREAD_STATE_WALK = 32,
    PYTHON_MAX_STACK_DEPTH = 128,
    PYTHON_MAX_FUNCNAME_LENGTH = 256,
    PYTHON_MAX_FILENAME_LENGTH = 256,
    MAX_PYTHON_SYMBOLS_SIZE = 200000,
    PYTHON_CFRAME_LINENO_ID = -1,
    PYTHON_UNSPECIFIED_OFFSET = -1,
};

enum python_frame_owner : u8 {
    FRAME_OWNED_BY_THREAD = 0,
    FRAME_OWNED_BY_GENERATOR = 1,
    FRAME_OWNED_BY_FRAME_OBJECT = 2,
    FRAME_OWNED_BY_CSTACK = 3,
};

struct python_ascii_object_offsets {
    u32 length_offset;
    u32 data_offset;
    u32 state_offset;
    u8 ascii_bit;
    u8 compact_bit;
    u8 statically_allocated_bit;
};

struct python_code_object_offsets {
    u32 co_firstlineno_offset;
    u32 filename_offset;
    u32 qualname_offset;
};

struct python_interpreter_frame_offsets {
    u32 f_code_offset;
    u32 previous_offset;
    u32 owner_offset;
};

struct python_cframe_offsets {
    u32 current_frame_offset;
};

struct python_internals_offsets {
    struct python_runtime_state_offsets py_runtime_state_offsets;
    struct python_thread_state_offsets py_thread_state_offsets;
    struct python_cframe_offsets py_cframe_offsets;
    struct python_interpreter_frame_offsets py_interpreter_frame_offsets;
    struct python_interpreter_state_offsets py_interpreter_state_offsets;
    struct python_code_object_offsets py_code_object_offsets;
    struct python_ascii_object_offsets py_ascii_object_offsets;
};

struct python_config {
    u64 py_thread_state_tls_offset;
    u64 py_runtime_relative_address;
    u32 version;

    struct python_internals_offsets offsets;
};

// hope this is enough to avoid collisions
// code_objects are usually allocated once,
//   so this should be good enough identifier within the process.
// Though add co_firstlineno which is quite granular
struct python_symbol_key {
    u64 code_object;
    u32 pid;
    int co_firstlineno;
};

struct python_symbol {
    char file_name[PYTHON_MAX_FILENAME_LENGTH];
    char qual_name[PYTHON_MAX_FUNCNAME_LENGTH];
};

struct python_code_object {
    u64 filename;
    u64 qualname;
};

struct python_frame {
    struct python_symbol_key symbol_key;
};

struct python_state {
    struct python_thread_key thread_key;
    struct python_config config;
    u64 py_runtime_address;
    struct python_frame frames[PYTHON_MAX_STACK_DEPTH];
    u32 frame_count;
    struct python_symbol symbol;
    struct python_symbol_key symbol_key;
    struct python_code_object code_object;
    u32 pid;
};

BPF_MAP(python_thread_id_py_thread_state, BPF_MAP_TYPE_LRU_HASH, struct python_thread_key, void*, MAX_PYTHON_THREADS);
BPF_MAP(python_symbols, BPF_MAP_TYPE_LRU_HASH, struct python_symbol_key, struct python_symbol, MAX_PYTHON_SYMBOLS_SIZE);
BPF_MAP(python_storage, BPF_MAP_TYPE_HASH, binary_id, struct python_config, MAX_BINARIES);
