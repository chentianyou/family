import os

import cv2


def main():
    images_path = "../../images"
    s_images_path = "../../simages"
    if not os.path.exists(s_images_path):
        os.mkdir(s_images_path)
    images = os.listdir(images_path)
    for image in images:
        input = images_path + "/" + image
        output = s_images_path + "/" + image
        im = cv2.imread(input)
        shape = im.shape
        reshape = (int(shape[1] / 5), int(shape[0] / 5))
        print(shape, reshape)
        im = cv2.resize(im, reshape)
        cv2.imwrite(output, im, [int(cv2.IMWRITE_JPEG_QUALITY), 90])
        cv2.imshow("family", im)
        cv2.waitKey(1)


if __name__ == '__main__':
    main()
