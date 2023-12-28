<template>
  <div class="tabs-view">
    <div class="tabs-view-main">
      <div ref="tabsWrap" class="tabs-line " :class="{'tabs-line-scroll--ed':state.scrollable}">
        <div class="tabs-line-left bg-#f5f7f9  text-#666 dark:text-#999" @click="scrollLeft" v-if="state.scrollable">
          <icon-line-md:chevron-left/>
        </div>
        <div class="tabs-line-right bg-#f5f7f9 text-#666 dark:text-#999" @click="scrollRight" v-if="state.scrollable">
          <icon-line-md:chevron-right/>
        </div>
        <div class="tabs-line-scroll " ref="tabsScroll">
          <!--     拖动     -->
          <Draggable :list="tabsList" animation="300" item-key="fullPath" class="flex">
            <template #item="{element}">
              <div class="tabs-line-scroll-item bg-#fff dark:bg-#333 flex items-center"
                   :id="`tab_item_${element.name}`" @click.stop="onTagClick(element)"
                   @contextmenu="onContextMenu($event, element)">
                <div class="h-32px flex items-center justify-center">
                  <span class="lh-14px text-14px "
                        :class="{'text-primary':state.activeTag==element.name}">{{ element.meta.title }}</span>
                  <icon-line-md:close class="text-14px ml-2px mr--6px  text-#999" v-if="!element.meta?.affix"
                                      @click.stop="onCloseTabs(element)"/>
                </div>
              </div>
            </template>
          </Draggable>
        </div>
      </div>
      <div class="tabs-close bg-#fff dark:bg-#333 text-#666 dark:text-#999">
        <n-dropdown trigger="hover" :options="tabsMenuOptions" placement="bottom-end" @select="onDropdownClick">
          <icon-solar:alt-arrow-down-linear/>
        </n-dropdown>
      </div>
      <n-dropdown
          :show="state.showDropdown"
          :x="state.dropdownX"
          :y="state.dropdownY"
          :options="tabsMenuOptions"
          @select="onDropdownClick"
          placement="bottom-start"
          @clickoutside="state.showDropdown=false"
      />
    </div>
  </div>
</template>
<script setup lang="ts">
import Draggable from 'vuedraggable'
import {computed, inject, nextTick, onMounted, reactive, ref, unref, watch} from "vue";
import {PageRoute} from "@/typings/route";
import {useRoute, useRouter} from "vue-router";
import elementResizeDetectorMaker from "element-resize-detector";
import {renderIcon} from "@/utils";
import {useRouteStore, useTabsStore} from "@/store";

const ut = useTabsStore()
const tabsList = computed(() => ut.tabList);

const router = useRouter();
const route = useRoute()
const tabsScroll = ref(null)
const tabsWrap = ref(null)
const state = reactive({
  activeTag: computed(() => route.name as string),
  dropdownX: 0,
  dropdownY: 0,
  showDropdown: false,
  scrollable: false,
  currentTab: null,
})
const pageIsAffix = ref(false)
const activePage = ref(null)
// 展开操作key
// 刷新 关闭 关闭其他 关闭所有
type ExpandKey = 'refresh' | 'close' | 'closeOther' | 'closeAll'
const route2PageRoute = (route: any): PageRoute => {
  return {
    name: route.name,
    path: route.path,
    meta: route.meta,
  }
}
const onDropdownClick = (key: ExpandKey) => {
  switch (key) {
    case 'refresh':
      refreshTabs()
      break;
    case 'close':
      closeCurrentTabs(state.currentTab ?? route2PageRoute(route))
      break;
    case 'closeOther':
      closeOtherTabs(route2PageRoute(route))
      break;
    case 'closeAll':
      closeAllTabs()
      break;
  }
  updateTabsScroll()
  state.showDropdown = false
}
const tabsMenuOptions = computed(() => {
  const isDisabled = tabsList.value.length <= 1
  let isRefresh = false
  if (activePage.value){
    isRefresh = activePage.value?.name != route.name
  }
  return [
    {
      label: '刷新页面',
      key: 'refresh',
      disabled: isRefresh,
      icon: renderIcon('line-md:backup-restore'),
    },
    {
      label: '关闭当前',
      key: 'close',
      disabled: isDisabled || pageIsAffix.value,
      icon: renderIcon('line-md:close-circle'),
    },
    {
      label: '关闭其他',
      key: 'closeOther',
      disabled: isDisabled,
      icon: renderIcon('line-md:close-circle'),
    },
    {
      label: '关闭所有',
      key: 'closeAll',
      icon: renderIcon('line-md:close-circle'),
    },
  ]

})
const onTagClick = (tag: PageRoute) => {
  router.push({name: tag.name})
}

// 右键菜单
const onContextMenu = (e: MouseEvent, tab: PageRoute) => {
  e.preventDefault();
  e.stopPropagation();
  activePage.value = tab
  pageIsAffix.value = tab.meta?.affix
  state.showDropdown = false;
  state.currentTab = tab;
  nextTick().then(() => {
    state.showDropdown = true;
    state.dropdownX = e.clientX;
    state.dropdownY = e.clientY;
  });
}
// 关闭
const onCloseTabs = (tag: PageRoute) => {
  closeCurrentTabs(tag)
}

const refreshTabs = async () => {
  const ur = useRouteStore()
  await ur.reloadPage()
  await updateTabsScroll()
}
const closeCurrentTabs = (page: PageRoute) => {
  if (page.meta?.affix) return
  ut.closeCurrentTab(page)
  // 关闭的当前页跳转到上一页
  if (page.name === state.activeTag) {
    const preTab = ut.tabList[Math.max(0, ut.tabList.length - 1)]
    router.push({name: preTab.name})
  }
  updateTabsScroll()
}
const closeOtherTabs = (page: PageRoute) => {
  ut.closeOtherTabs(page)
  updateTabsScroll()
}
const closeAllTabs = () => {
  ut.closeAllTabs()
  router.replace({path: '/'})
  updateTabsScroll()
}

async function updateTabsScroll(autoScroll = false) {
  await nextTick()
  if (!tabsScroll.value) return
  const wrapWidth = tabsScroll.value.scrollWidth
  const scrollWidth = tabsScroll.value.offsetWidth
  if (scrollWidth < wrapWidth) {
    if (autoScroll) {
      state.scrollable = true;
      let tagList = tabsScroll.value.querySelectorAll('.tabs-line-scroll-item') || [];
      [...tagList].forEach((tag: HTMLElement) => {
        // fix SyntaxError
        if (tag.id === `tab_item_${state.activeTag}`) {
          tag.scrollIntoView && tag.scrollIntoView();
        }
      });
    }
  } else {
    state.scrollable = false;
  }
}

/**
 * 滚动到指定位置
 * @param value 滚动的距离
 * @param amplitude 振幅
 */
function scrollTo(value: number, amplitude: number) {
  const currentScroll = tabsScroll.value.scrollLeft
  const scrollWidth =
      (amplitude > 0 && currentScroll + amplitude >= value) ||
      (amplitude < 0 && currentScroll + amplitude <= value)
          ? value
          : currentScroll + amplitude;
  tabsScroll.value && tabsScroll.value.scrollTo(scrollWidth, 0);
  if (scrollWidth === value) return;
  return window.requestAnimationFrame(() => scrollTo(value, amplitude));
}


/**
 * 向左滚动
 */
function scrollLeft() {
  const containerWidth = tabsScroll.value.offsetWidth;
  const currentScroll = tabsScroll.value.scrollLeft;
  if (!currentScroll) return;
  const scrollLeft = currentScroll > containerWidth ? currentScroll - containerWidth : 0;
  scrollTo(scrollLeft, (scrollLeft - currentScroll) / 20);
}

/**
 * 向右滚动
 */
function scrollRight() {
  const containerWidth = tabsScroll.value.offsetWidth;
  const navWidth = tabsScroll.value.scrollWidth;
  const currentScroll = tabsScroll.value.scrollLeft;
  if (navWidth - currentScroll <= containerWidth) return;
  const scrollLeft =
      navWidth - currentScroll > containerWidth * 2
          ? currentScroll + containerWidth
          : navWidth - containerWidth;
  scrollTo(scrollLeft, (scrollLeft - currentScroll) / 20);
}

/**
 * 监听窗口变化
 */
function onResize() {
  updateTabsScroll(true)
}

function onElementResize() {
  elementResizeDetectorMaker().listenTo(tabsScroll.value, onResize);
}

function onScroll(e) {
}

watch(() => route.name, (to) => {
  ut.addTab(route2PageRoute(route))
}, {immediate: true})

onMounted(() => {
  onElementResize()
})

window.addEventListener('scroll', onScroll, true)
</script>

<style scoped lang="scss">
.tabs-view {
  width: 100%;
  padding: 6px 0;
  display: flex;

  &-main {
    height: 32px;
    display: flex;
    max-width: 100%;
    min-width: 100%;
  }
}

.tabs-line {
  -webkit-box-flex: 1;
  flex-grow: 1;
  flex-shrink: 1;
  overflow: hidden;
  position: relative;

  &-scroll {
    white-space: nowrap;
    overflow: hidden;

    &--ed {
      padding: 0 32px;
      overflow: hidden;
    }

    &-item {
      height: 32px;
      padding: 0 12px;
      border-radius: 3px;
      margin-right: 6px;
      cursor: pointer;
      display: inline-block;
      position: relative;
      flex: 0 0 auto;
    }
  }

  .tabs-line-left, .tabs-line-right {
    width: 32px;
    text-align: center;
    position: absolute;
    height: 32px;
    cursor: pointer;
    font-size: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 3px;

    &:hover {
      background: rgba(0, 0, 0, 0.1);
    }
  }

  .tabs-line-left {
    left: 0;
  }

  .tabs-line-right {
    right: 0;
  }
}

.tabs-close {
  min-width: 32px;
  width: 32px;
  height: 32px;
  text-align: center;
  border-radius: 2px;
  cursor: pointer;
  font-size: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.tag-active-bg {
  background-color: var(--n-color)
}


</style>