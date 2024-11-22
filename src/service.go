package src


func CheckHasFit(ball Ball, box Box) bool {
	return ball.Diameter <= box.Height && ball.Diameter <= box.Width && ball.Diameter <= box.Length
}