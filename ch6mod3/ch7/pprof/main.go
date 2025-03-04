const TESTLENGTH = 20000
type DataType struct {
a,b,c,d,e,f,g int64
longByte []byte
}
func (dt DataType) init() {
}
var profile = flag.String("cpuprofile", "", "output pprof data to
file")
func main() {
flag.Parse()
if *profile != "" {
flag,err := os.Create(*profile)
if err != nil {
fmt.Println("Could not create profile",err)
}
pprof.StartCPUProfile(flag)
defer pprof.StopCPUProfile()
}
var wg sync.WaitGroup
numCPU := runtime.NumCPU()
runtime.GOMAXPROCS(numCPU)
wg.Add(TESTLENGTH)
for i := 0; i < TESTLENGTH; i++ {
go func() {
for y := 0; y < TESTLENGTH; y++ {
dT := DataType{}
dT.init()
}
wg.Done()
}()
}
wg.Wait()
fmt.Println("Complete.")
}