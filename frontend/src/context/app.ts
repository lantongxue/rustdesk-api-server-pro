import { useContext } from "@/hooks";
import { Ref, nextTick, ref } from "vue";

interface AppVueContext {
  isRouterAlive: Ref<boolean>
  reload: () => void;
}

const { useProvide: useAppVueProvide, useInject: useAppVueInject } = useContext<AppVueContext>();

export function useAppVueContext() {
  function reload() {
    context.isRouterAlive.value = false;
    nextTick(() => {
      context.isRouterAlive.value = true;
    });
  }

  const context: AppVueContext = {
    reload,
    isRouterAlive: ref(true)
  };

  function useProvide() {
    return useAppVueProvide(context);
  }

  return {
    useProvide,
    useInject: useAppVueInject,
  };
}
