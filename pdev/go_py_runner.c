#define Py_LIMITED_API
#include <Python.h>

PyObject * sum(PyObject *, PyObject *);

// Workaround missing variadic function support
// https://github.com/golang/go/issues/975
int PyArg_ParseTuple_LL(PyObject * args, long long * a, long long * b) {
    return PyArg_ParseTuple(args, "LL", a, b);
}

static PyMethodDef Go_py_runnerMethods[] = {
    {"sum", sum, METH_VARARGS, "Add two numbers."},
    {NULL, NULL, 0, NULL}
};

static struct PyModuleDef go_py_runnermodule = {
   PyModuleDef_HEAD_INIT, "go_py_runner", NULL, -1, Go_py_runnerMethods
};

PyMODINIT_FUNC
PyInit_go_py_runner(void)
{
    return PyModule_Create(&go_py_runnermodule);
}