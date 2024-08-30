import { View, Image } from "@tarojs/components";
import Taro, { useLoad } from "@tarojs/taro";

import desktop from "./desktop.png";
import { useMemo, useState } from "react";

export default function Index() {
  useLoad(() => {
    console.log("Page loaded.");
  });

  const srcQuery = Taro.getCurrentInstance().router?.params?.src!;
  const src = window.decodeURIComponent(srcQuery);

  const [type, setType] = useState(1); // 1 纯壁纸 2 桌面 3 锁屏

  const showType = (type: number) => () => {
    setType(type);
  }

  const download = () => {
    Taro.downloadFile({
      url: src,
      success: (res) => {
        if (res.statusCode === 200) {
          Taro.saveImageToPhotosAlbum({
            filePath: res.tempFilePath,
            success: () => {
              Taro.showToast({
                title: "保存成功",
                icon: "success",
              });
            },
          });
        }
      },
    });
  }

  const srcImage = useMemo(() => {
    if (type === 2) {
      return desktop
    } else {
      return desktop
    }
  }, [])

  return (
    <View className='box-border w-full h-screen mx-auto relative p-4'>
      <Image mode='aspectFit' className='w-full h-full absolute top-0 left-0 z-0' src={src} />

      {type > 1 && <Image mode='aspectFit' className='absolute top-0 left-0 right-0 bottom-0 w-full h-full z-10' src={srcImage} />}

      <View className='absolute bottom-4 left-4 right-4 h-16 bg-black bg-opacity-45 rounded-2xl z-20'>
        <View className='grid grid-cols-3 items-center justify-between h-full px-4'>
          <View className='text-white text-sm text-center block' onClick={showType(1)}>纯壁纸</View>
          <View className='text-white text-sm text-center block' onClick={showType(2)}>桌面</View>
          {/* <View className='text-white text-sm' onClick={showType(3)}>锁屏</View> */}
          <View className='text-white text-sm text-center block' onClick={download}>下载</View>
        </View>
      </View>

    </View>
  );
}
