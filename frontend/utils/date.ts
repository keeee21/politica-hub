import { ja } from 'date-fns/locale'
import { formatInTimeZone } from 'date-fns-tz'
import dayjs from 'dayjs'
import dayjsTimezone from 'dayjs/plugin/timezone.js'
import dayjsUTC from 'dayjs/plugin/utc.js'

const TZ = 'Asia/Tokyo'

dayjs.extend(dayjsTimezone)
dayjs.extend(dayjsUTC)

export function convISOStringToJST(date: string | Date): string {
  const d = typeof date === 'string' ? new Date(date) : date

  return formatInTimeZone(d, 'Asia/Tokyo', `yyyy-MM-dd HH:mm`)
}

export function convLocalDateTimeStringToJST(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy-MM-dd HH:mm:ss`)
}

export function convLocalDateTimeMinutesStringToJST(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy-MM-dd HH:mm`)
}

export function convLocalDateTimeStringToJSTSlash(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy/MM/dd HH:mm:ss`)
}

export function convLocalDateTimeStringToJSTSlashNoSeconds(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy/MM/dd HH:mm`)
}

export function convLocalDateStringToJST(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy-MM-dd`)
}

export function convLocalDateMonthStringToJST(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyyMM`)
}

export function convLocalDateMonthDayStringToJST(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyyMMdd`)
}

export function convLocalDateStringToJSTSlash(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy/MM/dd`)
}

export function convLocalDateMonthStringToJSTSlash(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy/MM`)
}

export function convLocalTimeStringToJST(date: Date): string {
  return formatInTimeZone(date, TZ, `HH:mm:ss`)
}

export function convFormatDate(date: Date): string {
  const dayName = ['日', '月', '火', '水', '木', '金', '土']
  const formattedDate
    = date.getFullYear()
      + '年'
      + (date.getMonth() + 1)
      + '月'
      + date.getDate()
      + '日'
      + '（'
      + dayName[date.getDay()]
      + '）'

  return formattedDate
}

// TODO:命名…
export function convForDisplayDayAndMonthAndinute(date: Date): string {
  return formatInTimeZone(date, TZ, `MM/dd(E)HH:mm`, { locale: ja })
}

// TODO:命名…
export function convForDisplayDayAndMonth(date: Date): string {
  return formatInTimeZone(date, TZ, `MM/dd(E)`, { locale: ja })
}

// TODO:命名…
export function convForDisplayHourAndMinute(date: Date): string {
  return formatInTimeZone(date, TZ, `HH:mm`)
}

// TODO:命名…
export function convForDisplayDayAndMonthOnly(date: Date): string {
  return formatInTimeZone(date, TZ, `MM/dd`, { locale: ja })
}

export function convLocalDateStringToJSTLocaled(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy年MM月`, { locale: ja })
}

export function convLocalDateStringToJSTLocaledYMD(date: Date): string {
  return formatInTimeZone(date, TZ, `yyyy年MM月dd日`, { locale: ja })
}

export function convForDisplayDate(date: Date, format: string): string {
  return formatInTimeZone(date, TZ, format, { locale: ja })
}

// TODO:命名…
export function convForDisplayOnlyDayOfWeek(date: Date): string {
  return formatInTimeZone(date, TZ, `(E)`, { locale: ja })
}

/**
 * Day.js オブジェクトを取得する（精製時の timezone 指定も可能）
 *
 * > let djs = dl.createDayjsObj(2022, 7, 1)
 * > djs.toJSON()
 * '2022-06-30T15:00:00.000Z'
 * > djs.toString()
 * 'Thu, 30 Jun 2022 15:00:00 GMT'
 * > djs.format()
 * '2022-07-01T00:00:00+09:00'
 *
 * @param yyyy
 * @param mm - Date と違い、そのまま月を渡すこと。
 * @param dd
 * @param h
 * @param m
 * @param s
 * @param tz
 */
export function createDayjsObj(
  yyyy: number,
  mm: number,
  dd: number,
  h = 0,
  m = 0,
  s = 0,
  tz?: string,
): dayjs.Dayjs {
  if (!tz) {
    tz = TZ
  }
  const ds = {
    yy: yyyy,
    mm: String(mm).padStart(2, '0'),
    dd: String(dd).padStart(2, '0'),
    h: String(h).padStart(2, '0'),
    m: String(m).padStart(2, '0'),
    s: String(s).padStart(2, '0'),
  }
  return dayjs.tz(`${ds.yy}-${ds.mm}-${ds.dd}T${ds.h}:${ds.m}:${ds.s}`, tz)
}

export function convertToDayjsObj(d: Date): dayjs.Dayjs {
  return dayjs(d).tz(TZ)
}

export function convertStringToDayjsObj(dateString: string): dayjs.Dayjs {
  return dayjs.tz(dateString, TZ)
}

/**
 *
 * @param yyyy
 * @param mm - Date と違い、そのまま月を渡すこと。
 * @param dd
 * @param h
 * @param m
 * @param s
 */
export function createLocalTimezoneDate(
  yyyy: number,
  mm: number,
  dd: number,
  h = 0,
  m = 0,
  s = 0,
): Date {
  // MEMO: 年月日は 0 で初期化
  // return new Date(yyyy, mm, dd, h, m, s); // これだと、process.env.TZ でしか操作できない
  /*
   * Javascript の Date() はインスタンスの中は UTC。
   * コンストラクタが受け取る日付は、 process.env.TZ とみなして受け取る。
   */

  return createDayjsObj(yyyy, mm, dd, h, m, s).toDate()
}

/**
 * 指定した日付が属する月の初日を Date 型で取得する
 *
 * @param yyyy
 * @param mm - Date と違い、そのまま月を渡すこと
 * @param dd
 * @param h
 * @param m
 * @param s
 */
export function startOfDayTZ(
  yyyy: number,
  mm: number,
  dd = 1,
  h = 0,
  m = 0,
  s = 0,
): Date {
  return createDayjsObj(yyyy, mm, dd, h, m, s).startOf('date').toDate()
}

/**
 * 指定した日付が属する月の初日を Date 型で取得する
 *
 * @param yyyy
 * @param mm - Date と違い、そのまま月を渡すこと
 * @param dd
 * @param h
 * @param m
 * @param s
 */
export function startOfMonthTZ(
  yyyy: number,
  mm: number,
  dd = 1,
  h = 0,
  m = 0,
  s = 0,
): Date {
  return createDayjsObj(yyyy, mm, dd, h, m, s).startOf('month').toDate()
}

/**
 * 指定した日付が属する月の末日日時を Date 型で取得する
 *
 * @param yyyy
 * @param mm - Date と違い、そのまま月を渡すこと
 * @param dd
 * @param h
 * @param m
 * @param s
 */
export function endOfMonthTZ(
  yyyy: number,
  mm: number,
  dd = 1,
  h = 0,
  m = 0,
  s = 0,
): Date {
  return createDayjsObj(yyyy, mm, dd, h, m, s).endOf('month').toDate()
}
