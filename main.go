package main

import (
	"github.com/FloatTech/gg"
)

func main() {
	const W, H = 512, 512
	canvas := gg.NewContext(W, H)

	canvas.SetRGBA255(163, 121, 134, 255)

	canvas.DrawEllipse(125+245/2, 76+230/2, 245/2, 230/2)
	canvas.Fill()

	canvas.DrawEllipse(129+222/2, 78+190/2, 222/2, 190/2)
	canvas.Fill()

	canvas.MoveTo(274.5, 298.5)
	canvas.QuadraticTo(311.5, 329.5, 328, 380)
	canvas.LineTo(366, 358)
	canvas.QuadraticTo(363, 348.5, 349.5, 331)
	canvas.QuadraticTo(364, 342, 372, 351)
	//
	canvas.LineTo(381, 339)
	canvas.QuadraticTo(355.5, 316, 274.5, 298.5)
	canvas.Fill()

	canvas.SetRGBA255(169, 223, 246, 255)

	canvas.MoveTo(274.5, 298.5)
	canvas.QuadraticTo(311.5, 329.5, 328, 380)
	canvas.LineTo(366, 358)
	canvas.QuadraticTo(363, 348.5, 349.5, 331)
	canvas.QuadraticTo(364, 342, 372, 351)
	//
	canvas.LineTo(375.5, 346)
	canvas.QuadraticTo(357.5, 325, 274.5, 298.5)
	canvas.Fill()

	canvas.SetRGBA255(254, 239, 214, 255)

	canvas.MoveTo(200, 135)
	canvas.QuadraticTo(187, 176, 165, 207)
	canvas.QuadraticTo(188, 300, 216, 337.5)
	canvas.QuadraticTo(260, 339, 292, 326)
	canvas.QuadraticTo(303, 315, 312, 300)
	canvas.LineTo(335, 233)
	canvas.QuadraticTo(317, 193, 294, 158)
	canvas.QuadraticTo(299, 182, 297, 205)
	canvas.QuadraticTo(284, 197, 269, 195)
	canvas.LineTo(261, 166)
	canvas.LineTo(263, 195)
	canvas.LineTo(257, 195)
	canvas.QuadraticTo(251, 169, 232, 129.5)
	canvas.QuadraticTo(225, 175.5, 216, 190.5)
	canvas.QuadraticTo(214.5, 162, 200, 135)
	canvas.Fill()

	canvas.MoveTo(322, 268)
	canvas.QuadraticTo(337.5, 264.5, 339, 279)
	canvas.QuadraticTo(340, 295, 325, 305)
	canvas.LineTo(314, 316)
	canvas.QuadraticTo(308, 317.5, 307, 310)
	canvas.Fill()

	canvas.SetRGBA255(227, 172, 165, 255)
	//耳朵
	canvas.MoveTo(322, 276)
	canvas.QuadraticTo(337.5, 274.5, 329, 289)
	canvas.QuadraticTo(326, 296, 312, 308)
	canvas.Fill()
	//嘴

	canvas.MoveTo(211, 286)
	canvas.CubicTo(192, 325, 258, 331, 254, 297)
	canvas.CubicTo(251, 288, 219, 281, 221, 286)
	canvas.Fill()

	canvas.SetRGBA255(163, 121, 134, 255)

	canvas.MoveTo(132, 228)
	canvas.QuadraticTo(151, 329, 169, 358)
	canvas.LineTo(206, 361)
	canvas.QuadraticTo(188, 335, 175, 305)
	canvas.QuadraticTo(196, 333, 219, 349)
	canvas.LineTo(227, 349)
	canvas.QuadraticTo(202, 321, 171, 247)
	canvas.LineTo(132, 228)
	canvas.Fill()

	canvas.MoveTo(332, 226)
	canvas.QuadraticTo(376, 304, 433, 349)
	canvas.QuadraticTo(437, 330, 444, 315)
	canvas.QuadraticTo(419, 289, 399, 261)
	canvas.QuadraticTo(428, 280, 460, 291)
	canvas.LineTo(466, 286)
	canvas.QuadraticTo(410, 262, 367, 197)
	canvas.LineTo(332, 226)
	canvas.Fill()

	canvas.SetRGBA255(169, 223, 246, 255)

	canvas.MoveTo(146, 232)
	canvas.QuadraticTo(150, 292, 204, 360)
	canvas.QuadraticTo(187, 333.5, 146, 232)
	canvas.Fill()

	canvas.MoveTo(359, 231)
	canvas.QuadraticTo(402, 290, 444, 315)
	canvas.LineTo(414, 283)
	canvas.LineTo(359, 231)
	canvas.Fill()

	canvas.SetRGBA255(238, 249, 253, 255)

	canvas.MoveTo(172, 233)
	canvas.QuadraticTo(185.5, 209, 218.5, 200.5)
	canvas.QuadraticTo(216, 234, 212, 249)
	canvas.QuadraticTo(186, 247, 172, 233)
	canvas.Fill()

	canvas.MoveTo(282.5, 219)
	canvas.QuadraticTo(273, 237.5, 270.5, 262.5)
	canvas.LineTo(293.5, 272)
	canvas.QuadraticTo(305.5, 269.5, 316, 262)
	canvas.QuadraticTo(305, 231.5, 282.5, 219)
	canvas.Fill()

	canvas.SetRGBA255(118, 83, 101, 255)

	canvas.MoveTo(166, 228)
	canvas.QuadraticTo(186.5, 204, 222, 200)
	canvas.LineTo(220, 207)
	canvas.QuadraticTo(188.5, 208, 170, 236)
	canvas.LineTo(166, 230)
	canvas.Fill()

	canvas.MoveTo(281.5, 215)
	canvas.QuadraticTo(310, 231, 319, 264)
	canvas.LineTo(314, 264)
	canvas.QuadraticTo(304.5, 234, 279.5, 220)
	canvas.LineTo(281.5, 215)
	canvas.Fill()

	canvas.MoveTo(186, 243)
	canvas.QuadraticTo(186, 227, 202, 206)
	canvas.QuadraticTo(209, 206, 216, 209)
	canvas.LineTo(212, 214)
	canvas.LineTo(215, 218.5)
	canvas.QuadraticTo(214, 236.5, 208, 249)
	canvas.LineTo(186, 243)
	canvas.Fill()

	canvas.MoveTo(273, 263)
	canvas.QuadraticTo(273, 248, 282.5, 226)
	canvas.QuadraticTo(292, 224.5, 298, 233)
	canvas.LineTo(295.5, 236.5)
	canvas.LineTo(300, 239.5)
	canvas.QuadraticTo(300, 258, 294, 271.5)
	canvas.LineTo(273, 263)
	canvas.Fill()

	canvas.SetRGBA255(163, 121, 134, 255)

	canvas.MoveTo(186.5, 242.5)
	canvas.QuadraticTo(187, 231, 196, 216)
	canvas.QuadraticTo(201, 223.5, 214, 222)
	canvas.QuadraticTo(213.5, 242, 208.5, 248)
	canvas.LineTo(186.5, 242.5)
	canvas.Fill()

	canvas.MoveTo(273, 262.5)
	canvas.QuadraticTo(273.5, 248, 277.5, 237)
	canvas.QuadraticTo(286, 243.5, 297, 241)
	canvas.QuadraticTo(300, 258.5, 294, 270.5)
	canvas.LineTo(273, 262.5)
	canvas.Fill()

	canvas.SetRGBA255(227, 172, 165, 255)

	canvas.MoveTo(186.5, 242.5)
	canvas.QuadraticTo(189, 239, 194, 233.5)
	canvas.QuadraticTo(200, 240, 207, 236.5)
	canvas.QuadraticTo(210, 241.5, 208.5, 248)
	canvas.LineTo(186.5, 242.5)
	canvas.Fill()

	canvas.MoveTo(273, 262)
	canvas.QuadraticTo(274, 256, 279, 252)
	canvas.QuadraticTo(282, 260, 294, 256)
	canvas.QuadraticTo(296, 263.5, 289, 269)
	canvas.MoveTo(274, 262)
	canvas.Fill()

	canvas.DrawEllipse(197+10/2, 218+18/2, 10/2, 18/2)
	canvas.Fill()

	canvas.DrawEllipse(282+10/2, 238+18/2, 10/2, 18/2)
	canvas.Fill()

	canvas.SetRGBA255(254, 230, 214, 255)

	canvas.MoveTo(190, 243)
	canvas.QuadraticTo(201, 238, 206, 248)
	canvas.LineTo(190, 243)
	canvas.Fill()

	canvas.MoveTo(276, 263)
	canvas.QuadraticTo(284, 258, 287, 268)
	canvas.LineTo(276, 263)
	canvas.Fill()

	canvas.SavePNG("./output.png")
}
