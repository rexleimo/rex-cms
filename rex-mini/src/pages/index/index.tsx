import { API } from "@/api";
import { View, Image } from "@tarojs/components";
import Taro, { useLoad } from "@tarojs/taro";
import { useState } from "react";


function ImageViewPort({ src }) {

  const toPreview = () => {
    Taro.navigateTo({
      url: `/pages/preview/index?src=${window.encodeURIComponent(src)}`
    })
  }

  return (
    <View onClick={toPreview} className='flex relative justify-center items-center rounded-md border border-solid border-gray-200'>
      <Image src={src} className='w-full rounded-md' />
      <View className='absolute box-border bottom-0 left-0 w-full bg-black/50 text-white p-1 text-[12px] rounded-md rounded-t-none'>
        <View className='text-center'>下载 0</View>
      </View>
    </View>
  )
}

export default function Index() {
  const [imageList, setImageList] = useState([]);
  useLoad(() => {
    console.log("Page loaded.");
    API.get("/");
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
