import { View, Image } from "@tarojs/components";
import { useLoad } from "@tarojs/taro";

const imageList = [
  "https://img0.baidu.com/it/u=1306108541,3890683509&fm=253&app=120&size=w931&n=0&f=JPEG&fmt=auto?sec=1723741200&t=9f34b107d0b6fdfacdcd27b974e7df0a",
  "https://img2.baidu.com/it/u=1490602810,3032519025&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1723741200&t=1fdebbc149e1a04430fc781ab6a50bfa",
  "https://raw.githubusercontent.com/rexleimo/rex-imgs/main/2024-08-08/1723048947119839900.jpg",
]

function ImageViewPort({ src }) {
  return (
    <View className='flex relative justify-center items-center rounded-md border border-solid border-gray-200'>
      <Image src={src} className='w-full rounded-md' />
      <View className='absolute box-border bottom-0 left-0 w-full bg-black/50 text-white p-1 text-[12px] rounded-md rounded-t-none'>
        <View className='text-center'>下载 0</View>
      </View>
    </View>
  )
}

export default function Index() {
  useLoad(() => {
    console.log("Page loaded.");
  });

  return (
    <View className='container mx-auto'>
      <View className='p-2'>
        <View className='flex justify-between items-center'>
          <View className='text-sm'>
            高级质感
          </View>
          <View className='text-xs text-slate-400'>探索壁纸美学之境</View>
        </View>
        <View className='grid gap-2 grid-cols-3 mt-2'>
          {
            imageList.map((src) => <ImageViewPort key={src} src={src} />)
          }
        </View>
      </View>
    </View>
  );
}
