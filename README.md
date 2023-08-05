# Video Converter
This microservice accepts video files, converts them to different qualities, and provides download links for original and converted versions. It is built using Golang, MongoDB, RabbitMQ and Kubernetes (k8s).

## Endpoints

- **POST /api/video**
  - Description: Uploads a video file and specifies the desired qualities for conversion.
  - Request Body:
    - file: The video file to be uploaded.
    - quality: Comma-separated list of qualities to convert the video to (e.g., "480,720").
  - Response: Returns a unique video ID for tracking the conversion progress and access to the original file and converted versions.

- **GET /api/video/info**
  - Description: Retrieves information about a specific video conversion.
  - Query Parameters:
    - id: The video ID obtained from the `/api/video` endpoint.
  - Response: Returns information about the video conversion, such as conversion status, original file/converted file id and size.

- **GET /video/VIDEOID**
  - Description: Retrieves the original video file.
  - Path Parameters:
    - VIDEOID: The video ID obtained from the `/api/video` endpoint.
  - Response: Returns the original video file.

- **GET /video/VIDEOID-QUALITY**
  - Description: Retrieves a specific converted version of the video file.
  - Path Parameters:
    - VIDEOID: The video ID obtained from the `/api/video` endpoint.
    - QUALITY: The file extension of the converted version (e.g., "480").
  - Response: Returns the converted version of the video file.

## Quality Conversion Support

The microservice supports the following qualities for video conversion:

- 144
- 240
- 360
- 480
- 720
- 1080
