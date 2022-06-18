const isReadyObj = () => {return (obj: {}): boolean => !!obj ? Object.keys(obj).length != 0 : false}
const isEmptyObj = () => {return (obj: {}): boolean => !!obj ? Object.keys(obj).length === 0 : true}
const isReadyArr = () => {return (arr: []): boolean => !!arr ? arr.length != 0 : false}
const isEmptyArr = () => {return (arr: []): boolean => !!arr ? arr.length === 0 : true}

const isNumber = () => {
  return (numVal: any): boolean | string => {
    const pattern = /^([1-9]\d*|0)(\.\d+)?$/;
    return pattern.test(numVal) || '半角数字のみです';
  }
}

const deepCopy = () => {
  return (obj: any): typeof obj => {
    return JSON.parse(JSON.stringify(obj))
  }
}

const changeToHalf = () => {
  return (str: string): undefined | string => {
    if(typeof str === 'number') return;
    return str.replace(/[Ａ-Ｚａ-ｚ０-９]/g, (s) => {
      return String.fromCharCode(s.charCodeAt(0) - 0xFEE0);
    });
  }
}

const changeToTimeStampFormat = () => {
  return (time: string): string => {
    const dateObj = new Date(time);
    const year = dateObj.getFullYear();
    const month = dateObj.getMonth() + 1;
    const date = dateObj.getDate();
    const dateStr = year + '/' + month + '/' + date;
    return dateStr;
  }
}

const changeToISOFormat = () => {
  return (time: string, dateSet = 0): string => {
    const dateObj = time ? new Date(time) : new Date();
    dateObj.setDate(dateObj.getDate() + dateSet);
    return dateObj.toISOString()
  }
}

const changeToDateISOFormat = () => {
  return (time: string, dateSet = 0): string => {
    const dateObj = time ? new Date(time) : new Date();
    dateObj.setDate(dateObj.getDate() + dateSet);
    const year = dateObj.getFullYear();
    const month = dateObj.getMonth() + 1;
    const date = dateObj.getDate();
    const dateStr = year + '-' + month + '-' + date;
    return dateStr;
  }
}

const resizeFile = async(e: File): Promise<string> => {
  const reader = await fileReader(e);
  const resizedImage = await resize(reader) as string;
  return resizedImage;
}

/** inputからバイト配列を読み込む */
const fileReader = (e: File): Promise<FileReader> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(e);
    reader.onload = () => {
      resolve(reader)
    }
    reader.onerror = reject
  })
}

/** base64からイメージを作成し、リサイズしてキャンバスに描画 */
const resize = (reader: FileReader, maxWidth = 900): Promise<string> => {
  return new Promise((resolve, reject) => {
    const image = new Image();
    image.src = reader.result as string;
    image.onload = () => {
      const imgType = image.src.substring(5, image.src.indexOf(';'));
      const imgWidth = maxWidth <= image.width ? maxWidth : image.width;
      const imgHeight = image.height * (imgWidth / image.width);
      const canvas = document.createElement('canvas');
      canvas.width = imgWidth;
      canvas.height = imgHeight;
      const ctx = canvas.getContext('2d') as CanvasRenderingContext2D;
      ctx.drawImage(image, 0, 0, imgWidth, imgHeight);
      const resizedImage = canvas.toDataURL(imgType)
      resolve(resizedImage)
    }
    image.onerror = reject
  })
}

const OK = 200;
const created = 201;
const accepted = 202;
const noContent = 204;
const notModified = 304;
const badRequest = 400;
const unauthorized = 401;
const forbidden = 403;
const notFound = 404;

const handleError = (status: number): void => {
  if(status === badRequest || status === unauthorized) return;
  return window.$nuxt.$router.back();
}
const checkStatus = () => {return (status: number, func: Function, error: Function = handleError) => {

  if(status === OK || status === created || status === accepted || status === noContent) {
    return func()
  }
  if(status === notModified) return window.$nuxt.$router.push('/login');
  if(status === badRequest || status === unauthorized || status === forbidden || status === notFound) return error(status);
}}

export { isReadyObj, isReadyArr, isEmptyArr, isEmptyObj, resizeFile, checkStatus, isNumber, changeToHalf, changeToISOFormat, changeToTimeStampFormat, changeToDateISOFormat, deepCopy}