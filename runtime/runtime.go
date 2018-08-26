// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, You can
// obtain one at http://mozilla.org/MPL/2.0/.

package runtime

import (
	"math"
	"os"

	"github.com/xtaniguchimasaya/iris/runtime/env"
	"github.com/xtaniguchimasaya/iris/runtime/ilos"
	"github.com/xtaniguchimasaya/iris/runtime/ilos/class"
	"github.com/xtaniguchimasaya/iris/runtime/ilos/instance"
)

func TopLevelHander(e env.Environment, c ilos.Instance) (ilos.Instance, ilos.Instance) {
	return nil, c
}

var TopLevel = env.NewEnvironment(
	instance.NewStream(os.Stdin, nil),
	instance.NewStream(nil, os.Stdout),
	instance.NewStream(nil, os.Stderr),
	instance.NewFunction(instance.NewSymbol("TOP-LEVEL-HANDLER"), TopLevelHander),
)

func defclass(name string, class ilos.Class) {
	symbol := instance.NewSymbol(name)
	TopLevel.Class.Define(symbol, class)
}

func defspecial(name string, function interface{}) {
	symbol := instance.NewSymbol(name)
	TopLevel.Special.Define(symbol, instance.NewFunction(func2symbol(function), function))
}

func defun(name string, function interface{}) {
	symbol := instance.NewSymbol(name)
	TopLevel.Function.Define(symbol, instance.NewFunction(symbol, function))
}

func defgeneric(name string, function interface{}) {
	symbol := instance.NewSymbol(name)
	lambdaList, _ := List(TopLevel, instance.NewSymbol("FIRST"), instance.NewSymbol("&REST"), instance.NewSymbol("REST"))
	generic := instance.NewGenericFunction(symbol, lambdaList, T, class.GenericFunction)
	generic.(*instance.GenericFunction).AddMethod(nil, lambdaList, []ilos.Class{class.StandardClass}, instance.NewFunction(symbol, function))
	TopLevel.Function.Define(symbol, generic)
}

func defglobal(name string, value ilos.Instance) {
	symbol := instance.NewSymbol(name)
	TopLevel.Variable.Define(symbol, value)
}

func init() {
	defglobal("*PI*", instance.Float(math.Pi))
	defglobal("*MOST-POSITIVE-FLOAT*", MostPositiveFloat)
	defglobal("*MOST-NEGATIVE-FLOAT*", MostNegativeFloat)
	defun("-", Substruct)
	defun("+", Add)
	defun("*", Multiply)
	defun("<", NumberLessThan)
	defun("<=", NumberLessThanOrEqual)
	defun("=", NumberEqual)
	defun(">", NumberGreaterThan)
	defun(">=", NumberGreaterThanOrEqual)
	defspecial("QUASIQUOTE", Quasiquote)
	defun("ABS", Abs)
	defspecial("AND", And)
	defun("APPEND", Append)
	defun("APPLY", Apply)
	defun("ARRAY-DIMENSIONS", ArrayDimensions)
	defun("AREF", Aref)
	defun("ASSOC", Assoc)
	// TODO: defspecial2("ASSURE", Assure)
	defun("ATAN", Atan)
	defun("ATAN2", Atan2)
	defun("ATANH", Atanh)
	defun("BASIC-ARRAY*-P", BasicArrayStarP)
	defun("BASIC-ARRAY-P", BasicArrayP)
	defun("BASIC-VECTOR-P", BasicVectorP)
	defspecial("BLOCK", Block)
	defun("CAR", Car)
	defspecial("CASE", Case)
	defspecial("CASE-USING", CaseUsing)
	defspecial("CATCH", Catch)
	defun("CDR", Cdr)
	defun("CEILING", Ceiling)
	defun("CERROR", Cerror)
	defun("CHAR-INDEX", CharIndex)
	defun("CHAR/=", CharNotEqual)
	defun("CHAR<", CharLessThan)
	defun("CHAR<=", CharLessThanOrEqual)
	defun("CHAR=", CharEqual)
	defun("CHAR>", CharGreaterThan)
	defun("CHAR>=", CharGreaterThanOrEqual)
	defun("CHARACTERP", Characterp)
	defspecial("CLASS", Class)
	defun("CLASS-OF", ClassOf)
	defun("CLOSE", Close)
	// TODO defun2("COERCION", Coercion)
	defspecial("COND", Cond)
	defun("CONDITION-CONTINUABLE", ConditionContinuable)
	defun("CONS", Cons)
	defun("CONSP", Consp)
	defun("CONTINUE-CONDITION", ContinueCondition)
	defspecial("CONVERT", Convert)
	defun("COS", Cos)
	defun("COSH", Cosh)
	defgeneric("CREATE", Create) //TODO Change to generic function
	defun("CREATE-ARRAY", CreateArray)
	defun("CREATE-LIST", CreateList)
	defun("CREATE-STRING", CreateString)
	defun("CREATE-STRING-INPUT-STREAM", CreateStringInputStream)
	defun("CREATE-STRING-OUTPUT-STREAM", CreateStringOutputStream)
	defun("CREATE-VECTOR", CreateVector)
	defspecial("DEFCLASS", Defclass)
	defspecial("DEFCONSTANT", Defconstant)
	defspecial("DEFDYNAMIC", Defdynamic)
	defspecial("DEFGENERIC", Defgeneric)
	defspecial("DEFMETHOD", Defmethod)
	defspecial("DEFGLOBAL", Defglobal)
	defspecial("DEFMACRO", Defmacro)
	defspecial("DEFUN", Defun)
	defun("DIV", Div)
	defspecial("DYNAMIC", Dynamic)
	defspecial("DYNAMIC-LET", DynamicLet)
	defun("ELT", Elt)
	defun("EQ", Eq)
	defun("EQL", Eql)
	defun("EQUAL", Equal)
	defun("ERROR", Error)
	defun("ERROR-OUTPUT", ErrorOutput)
	defun("EXP", Exp)
	defun("EXPT", Expt)
	// TODO defun2("FILE-LENGTH", FileLength)
	// TODO defun2("FILE-POSITION", FilePosition)
	// TODO defun2("FINISH-OUTPUT", FinishOutput)
	defspecial("FLET", Flet)
	defun("FLOAT", Float)
	defun("FLOATP", Floatp)
	defun("FLOOR", Floor)
	defspecial("FOR", For)
	defun("FORMAT", Format)
	defun("FORMAT-CHAR", FormatChar)
	defun("FORMAT-FLOAT", FormatFloat)
	defun("FORMAT-FRESH-LINE", FormatFreshLine)
	defun("FORMAT-INTEGER", FormatInteger)
	defun("FORMAT-OBJECT", FormatObject)
	defun("FORMAT-TAB", FormatTab)
	defun("FUNCALL", Funcall)
	defspecial("FUNCTION", Function)
	defun("FUNCTIONP", Functionp)
	defun("GAREF", Garef)
	defun("GCD", Gcd)
	defun("GENERAL-ARRAY*-P", GeneralArrayStarP)
	defun("GENERAL-VECTOR-P", GeneralVectorP)
	// TODO defun2("GENERIC-FUNCTION-P", GenericFunctionP)
	defun("GENSYM", Gensym)
	// TODO defun2("GET-INTERNAL-REAL-TIME", GetInternalRealTime)
	// TODO defun2("GET-INTERNAL-RUN-TIME", GetInternalRunTime)
	defun("GET-OUTPUT-STREAM-STRING", GetOutputStreamString)
	// TODO defun2("GET-UNIVERSAL-TIME", GetUniversalTime)
	defspecial("GO", Go)
	// TODO defun2("IDENTITY", Identity)
	defspecial("IF", If)
	// TODO defspecial2("IGNORE-ERRORS", IgnoreErrors)
	defgeneric("INITIALIZE-OBJECT", InitializeObject) // TODO change generic function
	defun("INPUT-STREAM-P", InputStreamP)
	defun("INSTANCEP", Instancep)
	// TODO defun2("INTEGER", Integer)
	defun("INTEGERP", Integerp)
	// TODO defun2("INTERNAL-TIME-UNITS-PER-SECOND", InternalTimeUnitsPerSecond)
	defun("ISQRT", Isqrt)
	defspecial("LABELS", Labels)
	defspecial("LAMBDA", Lambda)
	defun("LCM", Lcm)
	defun("LENGTH", Length)
	defspecial("LET", Let)
	defspecial("LET*", LetStar)
	defun("LIST", List)
	defun("LISTP", Listp)
	defun("LOG", Log)
	defun("MAP-INTO", MapInto)
	defun("MAPC", Mapc)
	defun("MAPCAN", Mapcan)
	defun("MAPCAR", Mapcar)
	defun("MAPCON", Mapcon)
	defun("MAPL", Mapl)
	defun("MAPLIST", Maplist)
	defun("MAX", Max)
	defun("MEMBER", Member)
	defun("MIN", Min)
	defun("MOD", Mod)
	defglobal("NI-L", Nil)
	defun("NOT", Not)
	defun("NREVERSE", Nreverse)
	defun("NULL", Null)
	defun("NUMBERP", Numberp)
	defun("OPEN-INPUT-FILE", OpenInputFile)
	defun("OPEN-IO-FILE", OpenIoFile)
	defun("OPEN-OUTPUT-FILE", OpenOutputFile)
	defun("OPEN-STREAM-P", OpenStreamP)
	defspecial("OR", Or)
	defun("OUTPUT-STREAM-P", OutputStreamP)
	defun("PARSE-NUMBER", ParseNumber)
	// TODO defun2("PREVIEW-CHAR", PreviewChar)
	// TODO defun2("PROVE-FILE", ProveFile)
	defspecial("PROGN", Progn)
	defun("PROPERTY", Property)
	defspecial("QUASIQUOTE", Quasiquote)
	defspecial("QUOTE", Quote)
	defun("QUOTIENT", Quotient)
	defun("READ", Read)
	// TODO defun2("READ-BYTE", ReadByte)
	defun("READ-CHAR", ReadChar)
	defun("READ-LINE", ReadLine)
	defun("REMOVE-PROPERTY", RemoveProperty)
	defun("REPORT-CONDITION", ReportCondition)
	defspecial("RETURN-FROM", ReturnFrom)
	defun("REVERSE", Reverse)
	defun("ROUND", Round)
	defun("SET-AREF", SetAref)
	defun("(SETF AREF)", SetAref)
	defun("SET-CAR", SetCar)
	defun("(SETF CAR)", SetCar)
	defun("SET-CDR", SetCdr)
	defun("(SETF CDR)", SetCdr)
	defun("SET-DYNAMIC", SetDynamic)
	defun("(SETF DYNAMIC)", SetDynamic)
	defun("SET-ELT", SetElt)
	defun("(SETF ELT)", SetElt)
	// TODO defun2("SET-FILE-POSITION", SetFilePosition)
	defun("SET-GAREF", SetGaref)
	defun("(SETF GAREF)", SetGaref)
	defun("SET-PROPERTY", SetProperty)
	defun("(SETF PROPERTY)", SetProperty)
	defspecial("SETF", Setf)
	defspecial("SETQ", Setq)
	defun("SIGNAL-CONDITION", SignalCondition)
	// TODO defun2("SIMPLE-ERROR-FORMAT-ARGUMENTS", SimpleErrorFormatArguments)
	// TODO defun2("SIMPLE-ERROR-FORMAT-STRING", SimpleErrorFormatString)
	defun("SIN", Sin)
	defun("SINH", Sinh)
	defun("SQRT", Sqrt)
	defun("STANDARD-INPUT", StandardInput)
	defun("STANDARD-OUTPUT", StandardOutput)
	defun("STREAM-READY-P", StreamReadyP)
	defun("STREAMP", Streamp)
	defun("STRING-APPEND", StringAppend)
	defun("STRING-INDEX", StringIndex)
	defun("STRING/=", StringNotEqual)
	defun("STRING>", StringGreaterThan)
	defun("STRING>=", StringGreaterThanOrEqual)
	defun("STRING=", StringEqual)
	defun("STRING<", StringLessThan)
	defun("STRING<=", StringLessThanOrEqual)
	defun("STRINGP", Stringp)
	defun("SUBCLASSP", Subclassp)
	defun("SUBSEQ", Subseq)
	defun("SYMBOLP", Symbolp)
	defglobal("T", T)
	defspecial("TAGBODY", Tagbody)
	defspecial("TAN", Tan)
	defspecial("TANH", Tanh)
	// TODO defspecial2("THE", The)
	defspecial("THROW", Throw)
	defun("TRUNCATE", Truncate)
	// TODO defun1("UNDEFINED-ENTITY-NAME", UndefinedEntityName)
	// TODO defun2("UNDEFINED-ENTITY-NAMESPACE", UndefinedEntityNamespace)
	defspecial("UNWIND-PROTECT", UnwindProtect)
	defun("VECTOR", Vector)
	defspecial("WHILE", While)
	defspecial("WITH-ERROR-OUTPUT", WithErrorOutput)
	defspecial("WITH-HANDLER", WithHandler)
	defspecial("WITH-OPEN-INPUT-FILE", WithOpenInputFile)
	defspecial("WITH-OPEN-OUTPUT-FILE", WithOpenOutputFile)
	defspecial("WITH-STANDARD-INPUT", WithStandardInput)
	defspecial("WITH-STANDARD-OUTPUT", WithStandardOutput)
	// TODO defun2("WRITE-BYTE", WriteByte)

	defclass("<OBJECT>", class.Object)
	defclass("<BUILT-IN-CLASS>", class.BuiltInClass)
	defclass("<STANDARD-CLASS>", class.StandardClass)
	defclass("<BASIC-ARRAY>", class.BasicArray)
	defclass("<BASIC-ARRAY-STAR>", class.BasicArrayStar)
	defclass("<GENERAL-ARRAY-STAR>", class.GeneralArrayStar)
	defclass("<BASIC-VECTOR>", class.BasicVector)
	defclass("<GENERAL-VECTOR>", class.GeneralVector)
	defclass("<STRING>", class.String)
	defclass("<CHARACTER>", class.Character)
	defclass("<FUNCTION>", class.Function)
	defclass("<GENERIC-FUNCTION>", class.GenericFunction)
	defclass("<STANDARD-GENERIC-FUNCTION>", class.StandardGenericFunction)
	defclass("<LIST>", class.List)
	defclass("<CONS>", class.Cons)
	defclass("<NULL>", class.Null)
	defclass("<SYMBOL>", class.Symbol)
	defclass("<NUMBER>", class.Number)
	defclass("<INTEGER>", class.Integer)
	defclass("<FLOAT>", class.Float)
	defclass("<SERIOUS-CONDITION>", class.SeriousCondition)
	defclass("<ERROR>", class.Error)
	defclass("<ARITHMETIC-ERROR>", class.ArithmeticError)
	defclass("<DIVISION-BY-ZERO>", class.DivisionByZero)
	defclass("<FLOATING-POINT-ONDERFLOW>", class.FloatingPointOnderflow)
	defclass("<FLOATING-POINT-UNDERFLOW>", class.FloatingPointUnderflow)
	defclass("<CONTROL-ERROR>", class.ControlError)
	defclass("<PARSE-ERROR>", class.ParseError)
	defclass("<PROGRAM-ERROR>", class.ProgramError)
	defclass("<DOMAIN-ERROR>", class.DomainError)
	defclass("<UNDEFINED-ENTITY>", class.UndefinedEntity)
	defclass("<UNDEFINED-VARIABLE>", class.UndefinedVariable)
	defclass("<UNDEFINED-FUNCTION>", class.UndefinedFunction)
	defclass("<SIMPLE-ERROR>", class.SimpleError)
	defclass("<STREAM-ERROR>", class.StreamError)
	defclass("<END-OF-STREAM>", class.EndOfStream)
	defclass("<STORAGE-EXHAUSTED>", class.StorageExhausted)
	defclass("<STANDARD-OBJECT>", class.StandardObject)
	defclass("<STREAM>", class.Stream)
}
