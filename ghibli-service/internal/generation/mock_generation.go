package generation

import "fmt"

type GhibliImage struct {
    URL string `json:"url"`
}

func GenerateMock(imageID string, count int) []GhibliImage {
    var results []GhibliImage
    for i := 0; i < count; i++ {
        results = append(results, GhibliImage{
            URL: fmt.Sprintf("https://ghibli-service/fake-ghibli/%s_%d.png", imageID, i),
        })
    }
    return results
}
