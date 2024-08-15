import { View, Image } from "@tarojs/components";
import Taro, { useLoad } from "@tarojs/taro";

import desktop from "./desktop.png";

export default function Index() {
  useLoad(() => {
    console.log("Page loaded.");
  });

  const srcQuery = Taro.getCurrentInstance().router?.params?.src!;
  const src = window.decodeURIComponent(srcQuery);

  return (
    <View className='box-border w-full h-full mx-auto relative p-4'>
      <Image mode='aspectFit' className='w-full h-full absolute top-0 left-0' src={src} />
      <Image mode='aspectFit' className='w-full h-full ' src={desktop} />

      <View className='absolute bottom-4 left-4 right-4 h-16 bg-black bg-opacity-45 rounded-2xl'>
        <View className='flex items-center justify-between h-full px-4'>
          <View className='text-white text-sm'>纯壁纸</View>
          <View className='text-white text-sm'>桌面</View>
          <View className='text-white text-sm'>锁屏</View>
          <View className='text-white text-sm'>下载</View>
        </View>
      </View>

    </View>
  );
}
