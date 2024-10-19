package myLib 
import "testing"
func TestHello(t *testing.T){
  got:= hello()
  want="hello world"
  if got!= want {
    t.Errorf("got %q want %q",got, want)
    }
  
  }
func TestName(t *testing.T){
  got:=name()
  want:= "pawan"
  
  if got != want{
    t.Errorf("got %q want %q ",git,want)
    }
}
func Testschool(t*testing.T){
  got:=school()
  want:="gvpce"
  if gor!= want {
    t.Errorf("got %q want %q",got, want)
    }
  }
