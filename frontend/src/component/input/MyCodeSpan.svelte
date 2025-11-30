<script lang="ts">
    import { dark_mode } from "../../store/changeBody";
    export let style_in = "";
</script>
<span class={"code" + ($dark_mode ? ' code-dark' : '')} style={style_in}><slot /></span>
<style>
.code {
    box-shadow: 0 0 3px black;
    background-color: lightgray;
    border-radius: 4px;
    color: black;
    padding: 0 5px;
    user-select: all;
    -webkit-user-select: all;
    -moz-user-select: all;
    -ms-user-select: all;
    font-size: 12px;
}
.code-dark {
    box-shadow: 0 0 3px white;
    background-color: rgb(44, 44, 44);
    color: white;
}
</style>