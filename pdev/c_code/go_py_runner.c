#define Py_LIMITED_API
#include <Python.h>

static PyObject *
sum(PyObject *self, PyObject *args)
{
    const long long a, b;

    if (!PyArg_ParseTuple(args, "LL", &a, &b))
        return NULL;

    return PyLong_FromLongLong(a + b);
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