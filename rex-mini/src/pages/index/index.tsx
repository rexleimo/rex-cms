import { View, Image } from "@tarojs/components";
import { useLoad } from "@tarojs/taro";

const imageList = [
  "https://img0.baidu.com/it/u=1306108541,3890683509&fm=253&app=120&size=w931&n=0&f=JPEG&fmt=auto?sec=1723741200&t=9f34b107d0b6fdfacdcd27b974e7df0a",
  "https://img2.baidu.com/it/u=1490602810,3032519025&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1723741200&t=1fdebbc149e1a04430fc781ab6a50bfa"
]

function ImageViewPort({ src }) {
  return (
    <View className='flex justify-center items-center'>
      <Image src={src} className='w-full' />
    </View>
  )
}

export default function Index() {
  useLoad(() => {
    console.log("Page loaded.");
  });

  return (
    <View className='container'>
      <View className='p-3'>
        <View className='grid gap-3 grid-cols-3'>
          {
            imageList.map((src) => <ImageViewPort key={src} src={src} />)
          }
        </View>
      </View>
    </View>
  );
}
